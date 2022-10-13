package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn/publisher-kafka/usecase"
)

type PublishHandler struct {
	MessageUsecase usecase.MessageUsecase
}

func NewPublishHandler(messageUsecase usecase.MessageUsecase) PublishHandler {
	return PublishHandler{MessageUsecase: messageUsecase}
}

func (ch PublishHandler) InitRouter(router *gin.Engine) {
	router.GET("/publish", ch.Publish)
}

func (ch PublishHandler) Publish(c *gin.Context) {
	err := ch.MessageUsecase.Publish()
	if err != nil {
		log.Printf("[handler][Publish] error while do Publish %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error", "InternalMessage": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Success", "Response": "success"})
	c.Abort()
}
