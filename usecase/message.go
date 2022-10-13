package usecase

import (
	"fmt"
	"log"

	"github.com/learn/publisher-kafka/repository"
)

type messageUsecase struct {
	MessageRepo repository.MessageRepository
}

func NewMessageUsecase(messageRepo repository.MessageRepository) MessageUsecase {
	return messageUsecase{
		MessageRepo: messageRepo,
	}
}

func (c messageUsecase) Publish() (err error) {
	counter := 0
	incrementPublish := 20
	for i := 0; i < incrementPublish; i++ {
		if i%2 == 0 {
			counter++
		}

		key := fmt.Sprintf("key-A%d", counter)
		err = c.MessageRepo.Publish("testingMultiple", key, []byte(fmt.Sprintf("%s-message-%d", key, i)))
		if err != nil {
			log.Println("usecase publish error: ", err)
			return err
		}

	}

	return nil
}

type MessageUsecase interface {
	Publish() (err error)
}
