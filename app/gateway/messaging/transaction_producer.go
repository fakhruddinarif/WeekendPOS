package messaging

import (
	"WeekendPOS/app/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type TransactionProducer struct {
	Producer[*model.TransactionEvent]
}

func NewTransactionProducer(producer *kafka.Producer, log *logrus.Logger) *TransactionProducer {
	return &TransactionProducer{
		Producer: Producer[*model.TransactionEvent]{
			Producer: producer,
			Topic:    "transactions",
			Log:      log,
		},
	}
}
