package messaging

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type ProductConsumer struct {
	Log *logrus.Logger
}

func NewProductConsumer(log *logrus.Logger) *ProductConsumer {
	return &ProductConsumer{
		Log: log,
	}
}

func (c *ProductConsumer) Consume(message *kafka.Message) error {
	productEvent := new(model.ProductEvent)
	if err := json.Unmarshal(message.Value, productEvent); err != nil {
		c.Log.Errorf("Failed unmarshal product event : %+v", err)
		return err
	}
	c.Log.Infof("Received topic addresses with event: %v from partition %d", productEvent, message.TopicPartition.Partition)
	return nil
}
