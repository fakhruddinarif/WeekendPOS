package messaging

import (
	"WeekendPOS/app/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Producer[T model.Event] struct {
	Producer *kafka.Producer
	Topic    string
	Log      *logrus.Logger
}

func (p *Producer[T]) GetTopic() *string {
	return &p.Topic
}

func (p *Producer[T]) Send(event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     p.GetTopic(),
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(event.GetId()),
	}

	err = p.Producer.Produce(message, nil)
	if err != nil {
		p.Log.WithError(err).Error("failed to produce message")
		return err
	}

	return nil
}
