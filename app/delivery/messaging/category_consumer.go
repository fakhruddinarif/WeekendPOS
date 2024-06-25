package messaging

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type CategoryConsumer struct {
	Log *logrus.Logger
}

func NewCategoryConsumer(log *logrus.Logger) *CategoryConsumer {
	return &CategoryConsumer{Log: log}
}

func (c *CategoryConsumer) Consume(message *kafka.Message) error {
	categoryEvent := new(model.CategoryEvent)
	if err := json.Unmarshal(message.Value, categoryEvent); err != nil {
		c.Log.WithError(err).Error("Failed to unmarshal category event")
		return err
	}
	c.Log.Infof("Received topic contacts with event: %v from partition %d", categoryEvent, message.TopicPartition.Partition)
	return nil
}
