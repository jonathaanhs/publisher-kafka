package repository

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type messageRepo struct {
	KafkaProducer *kafka.Producer
}

func NewMessageRepository(kafkaProducer *kafka.Producer) MessageRepository {
	return messageRepo{
		KafkaProducer: kafkaProducer,
	}
}

func (kr messageRepo) Publish(topic string, key string, message []byte) (err error) {
	deliverChan := make(chan kafka.Event)

	go func() {
		err := kr.KafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(key),
			Value: message,
		}, deliverChan)
		if err != nil {
			log.Println("Error while produce: ", err)
			deliverChan <- nil
		}
	}()

	kafkEvent := <-deliverChan
	if err != nil {
		return err
	}

	msg := kafkEvent.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		log.Println("Error while publish: ", msg.TopicPartition.Error)
		return err
	}

	return err
}

type MessageRepository interface {
	Publish(topic string, key string, message []byte) (err error)
}
