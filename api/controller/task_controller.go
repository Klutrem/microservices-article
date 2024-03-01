package controller

import (
	"encoding/json"
	"fmt"

	domain_common "main/domain/domainCommon"
	"main/lib"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	kafkaService lib.KafkaClient
}

func NewTaskController(kafka lib.KafkaClient) TaskController {
	return TaskController{
		kafkaService: kafka,
	}
}
func (tc *TaskController) TestReplyTopic(c *gin.Context) {
	message, err := json.Marshal("test_message")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	response, err := tc.kafkaService.SendWithReply(lib.TestReplyTopic, message)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, string(response))
}

type dto struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (tc *TaskController) TestConsumeTopic(message domain_common.KafkaMessage) (response []byte, err error) {
	println("slepping for 5 seconds", string(message.Value))
	resp := dto{
		Name:    "test",
		Surname: "test",
	}

	response, err = json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (tc *TaskController) TestSecondTopic(replyTopic string, message []byte) {
	fmt.Println("getting message from second topic:", string(message))
}
