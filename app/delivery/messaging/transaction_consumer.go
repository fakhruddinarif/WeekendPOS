package messaging

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type TransactionConsumer struct {
	Log *logrus.Logger
}

func NewTransactionConsumer(log *logrus.Logger) *TransactionConsumer {
	return &TransactionConsumer{
		Log: log,
	}
}

func (c *TransactionConsumer) Consume(message *kafka.Message) error {
	transactionEvent := new(model.TransactionEvent)
	if err := json.Unmarshal(message.Value, transactionEvent); err != nil {
		c.Log.Errorf("Failed unmarshal transaction event : %+v", err)
		return err
	}
	c.Log.Infof("Received topic addresses with event: %v from partition %d", transactionEvent, message.TopicPartition.Partition)
	return nil
}
