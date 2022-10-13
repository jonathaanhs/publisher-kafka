package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/learn/publisher-kafka/delivery"
	"github.com/learn/publisher-kafka/repository"
	"github.com/learn/publisher-kafka/usecase"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("broker"),
	})
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	messageRepo := repository.NewMessageRepository(kafkaProducer)
	messageUsecase := usecase.NewMessageUsecase(messageRepo)

	publishHandler := delivery.NewPublishHandler(messageUsecase)

	publishHandler.InitRouter(router)

	log.Println("Starting HTTP server")

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
