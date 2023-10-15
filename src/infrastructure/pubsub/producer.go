package pubsub

import (
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

const (
	defaultRetries  = 3
	defaultLingerMs = 5
)

type Producer interface {
	Publish(topic string, data []byte, headers []kafka.Header) error
	PublishWithKey(topic string, key string, data []byte, headers []kafka.Header) error
	Close()
}

type producerImpl struct {
	producer *kafka.Producer
	wg       sync.WaitGroup
	logger   *zap.Logger
}

func NewProducer(serversAddr, clientId string, logger *zap.Logger) (Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"client.id":         clientId,
		"bootstrap.servers": serversAddr,
		"retries":           defaultRetries,
		"linger.ms":         defaultLingerMs,
	})
	if err != nil {
		return nil, err
	}

	producer := producerImpl{
		producer: p,
		logger:   logger.Named("kafka_producer"),
	}

	producer.wg.Add(1)
	go producer.reportHandlerLoop()

	return &producer, nil
}

func (p *producerImpl) reportHandlerLoop() {
	defer p.wg.Done()

	for e := range p.producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			m := ev
			if m.TopicPartition.Error != nil {
				p.logger.Warn("failed to deliver message to kafka topic",
					zap.Error(m.TopicPartition.Error),
					zap.String("topic", *m.TopicPartition.Topic),
				)
			} else {
				p.logger.Debug("delivered message to kafka topic",
					zap.String("topic", *m.TopicPartition.Topic),
					zap.Int32("partition", m.TopicPartition.Partition),
					zap.Any("offset", m.TopicPartition.Offset),
				)
			}
		default:
			p.logger.Debug("ignored event", zap.String("event", ev.String()))
		}
	}
}

func (p *producerImpl) Publish(topic string, data []byte, headers []kafka.Header) error {
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     data,
		Headers:   headers,
		Timestamp: time.Now().UTC(),
	}

	return p.publish(msg)
}

func (p *producerImpl) PublishWithKey(topic string, key string, data []byte, headers []kafka.Header) error {
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     data,
		Headers:   headers,
		Key:       []byte(key),
		Timestamp: time.Now().UTC(),
	}
	return p.publish(msg)
}

func (p *producerImpl) publish(msg kafka.Message) error {
	if err := p.producer.Produce(&msg, nil); err != nil {
		p.logger.Error("error publishing message to kafka topic",
			zap.Error(err),
			zap.String("topic", *msg.TopicPartition.Topic),
		)
		return err
	}
	return nil
}

func (p *producerImpl) Close() {
	p.producer.Flush(1000)
	p.producer.Close()
	p.wg.Wait()
}
