package messaging

import (
	"WeekendPOS/app/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type ProductProducer struct {
	Producer[*model.ProductEvent]
}

func NewProductProducer(producer *kafka.Producer, log *logrus.Logger) *ProductProducer {
	return &ProductProducer{
		Producer: Producer[*model.ProductEvent]{
			Producer: producer,
			Topic:    "products",
			Log:      log,
		},
	}
}
