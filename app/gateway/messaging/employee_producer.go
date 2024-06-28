package messaging

import (
	"WeekendPOS/app/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type EmployeeProducer struct {
	Producer[*model.EmployeeEvent]
}

func NewEmployeeProducer(producer *kafka.Producer, log *logrus.Logger) *EmployeeProducer {
	return &EmployeeProducer{
		Producer: Producer[*model.EmployeeEvent]{
			Producer: producer,
			Topic:    "employees",
			Log:      log,
		},
	}
}
