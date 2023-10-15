package integrationevents

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/Azamjon99/eater-service/src/infrastructure/config"
	"github.com/Azamjon99/eater-service/src/infrastructure/pubsub"
	"go.uber.org/zap"
)

var (
	topics = []string{
		"order.created",
		"order.canceled",
		"card.added",
		"card.deleted",
		"restaurant.rated",
		"restaurant_rating.updated",
		"order.rated",
		"order_rating.updated",
	}
)

const (
	clientID = "eater-svc-producer"
)

func NewProducer(ctx context.Context, conf config.Config, logger *zap.Logger) pubsub.Producer {
	// create topics
	if err := createTopics(ctx, conf.KafkaAddress, logger); err != nil {
		logger.Warn("error creating topics", zap.Error(err))
	}

	// create producer
	producer, err := pubsub.NewProducer(conf.KafkaAddress, clientID, logger)
	if err != nil {
		log.Fatalf("Error creating pubsub.Producer: %v", err)
	}
	return producer
}

func createTopics(ctx context.Context, serverAddr string, logger *zap.Logger) error {
	adminConfig := &kafka.ConfigMap{
		"bootstrap.servers": serverAddr,
	}

	// Create an AdminClient
	adminClient, err := kafka.NewAdminClient(adminConfig)
	if err != nil {
		return fmt.Errorf("error creating AdminClient: %s", err)
	}
	defer adminClient.Close()

	var specs []kafka.TopicSpecification

	for _, topic := range topics {
		specs = append(specs, kafka.TopicSpecification{
			Topic:             topic,
			NumPartitions:     1, // Number of partitions for the topic
			ReplicationFactor: 1, // Replication factor for the topic
		})
	}

	// Create the topic
	results, err := adminClient.CreateTopics(
		ctx,
		specs,
		kafka.SetAdminOperationTimeout(5*time.Second),
	)
	if err != nil {
		return fmt.Errorf("error creating topic: %s", err)
	}

	// Print the results
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			if result.Error.Code() != kafka.ErrNoError {
				return fmt.Errorf("failed to create topic: %s", result.Error)
			}
		}
	}

	return nil
}