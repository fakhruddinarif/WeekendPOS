package messaging

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type EmployeeConsumer struct {
	Log *logrus.Logger
}

func NewEmployeeConsumer(log *logrus.Logger) *EmployeeConsumer {
	return &EmployeeConsumer{
		Log: log,
	}
}

func (c *EmployeeConsumer) Consume(message *kafka.Message) error {
	employee := new(model.EmployeeEvent)
	if err := json.Unmarshal(message.Value, employee); err != nil {
		c.Log.Errorf("Failed unmarshal employee event : %+v", err)
		return err
	}
	c.Log.Infof("Received topic addresses with event: %v from partition %d", employee, message.TopicPartition.Partition)
	return nil
}
