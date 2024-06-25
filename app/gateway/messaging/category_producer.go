package messaging

import (
	"WeekendPOS/app/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type CategoryProducer struct {
	Producer[*model.CategoryEvent]
}

func NewCategoryProducer(producer *kafka.Producer, log *logrus.Logger) *CategoryProducer {
	return &CategoryProducer{
		Producer: Producer[*model.CategoryEvent]{
			Producer: producer,
			Topic:    "categories",
			Log:      log,
		},
	}
}
