package pubsub

import (
	"context"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type Consumer interface {
	Subscribe(topic string, handler MessageHandler) error
}

type MessageHandler func(msg *kafka.Message) error

type ConsumerOption func(c *consumerImpl)

type consumerImpl struct {
	ctx      context.Context
	consumer *kafka.Consumer
	wg       sync.WaitGroup
	topics   []string
	handlers map[string]MessageHandler
	mu       sync.RWMutex
	logger   *zap.Logger
}

func NewConsumer(
	ctx context.Context,
	serversAddr, clientId, groupId string,
	options ...ConsumerOption,
) (Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"client.id":          clientId,
		"group.id":           groupId,
		"bootstrap.servers":  serversAddr,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
		"session.timeout.ms": 6000,
	})
	if err != nil {
		return nil, err
	}

	consumer := consumerImpl{
		ctx:      ctx,
		consumer: c,
		topics:   make([]string, 0),
		handlers: make(map[string]MessageHandler),
	}

	for _, opt := range options {
		opt(&consumer)
	}

	go consumer.waitForCancel()
	consumer.wg.Add(1)
	go consumer.startPoll()
	return &consumer, nil
}

func (c *consumerImpl) Subscribe(topic string, handler MessageHandler) error {
	if topic == "" {
		return nil
	}

	c.logger.Debug("subscribing to topic", zap.String("topic", topic))
	c.addHandler(topic, handler)
	if err := c.consumer.SubscribeTopics(c.topics, nil); err != nil {
		return err
	}

	return nil
}

func (c *consumerImpl) startPoll() {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("stopping poll loop")
			return
		default:
			ev := c.consumer.Poll(1000)
			if ev == nil {
				// if there are no messages, polling will timeout and will return nil event
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				c.logger.Debug("received event", zap.Any("message", e))
				c.handleMsg(e)
			case kafka.Error:
				c.logger.Error("received error event", zap.Error(e))
			case kafka.AssignedPartitions:
				c.logger.Debug("new partition(s) assigned", zap.Any("event", e))
				c.consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				c.logger.Debug("partition(s) revoked", zap.Any("event", e))
				commitedOffsets, err := c.consumer.Commit()
				if err != nil && err.(kafka.Error).Code() != kafka.ErrNoOffset {
					c.logger.Error("failed to commit offsets", zap.Error(err))
					continue
				}

				c.logger.Debug("commited offset", zap.Any("commited offsets", commitedOffsets))
				c.consumer.Unassign()
			default:
				c.logger.Debug("unexpected event type", zap.Any("event", e))
			}
		}
	}
}

func (c *consumerImpl) waitForCancel() {
	<-c.ctx.Done()
	c.wg.Wait()
	c.consumer.Unsubscribe()
	c.consumer.Close()
}

func (c *consumerImpl) handleMsg(msg *kafka.Message) {
	handler := c.getHandler(*msg.TopicPartition.Topic)
	if handler == nil {
		c.logger.Warn("handler not found for the given topic", zap.String("topic", *msg.TopicPartition.Topic))
		return
	}

	if err := handler(msg); err != nil {
		c.logger.Error("error handling msg",
			zap.Error(err),
		)
		return
	}

	_, err := c.consumer.CommitMessage(msg)
	if err != nil {
		c.logger.Warn("error commiting message", zap.Error(err))
	}
}

func (c *consumerImpl) getHandler(topic string) MessageHandler {
	c.mu.RLock()
	defer c.mu.RUnlock()

	handler, ok := c.handlers[topic]
	if !ok {
		return nil
	}

	return handler
}

func (c *consumerImpl) addHandler(topic string, handler MessageHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.handlers[topic]; !ok {
		c.topics = append(c.topics, topic)
	}
	c.handlers[topic] = handler
}

func WithLogger(logger *zap.Logger) ConsumerOption {
	return func(c *consumerImpl) {
		c.logger = logger.Named("kafka_consumer")
	}
}
