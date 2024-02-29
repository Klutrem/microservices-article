package controller

import (
	"encoding/json"
	"fmt"

	"main/lib"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	KafkaClient lib.KafkaClient
}

func NewTaskController(kafka lib.KafkaClient) TaskController {
	return TaskController{
		KafkaClient: kafka,
	}
}
func (tc *TaskController) TestReplyTopic(c *gin.Context) {
	message, err := json.Marshal("test_message")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	response, err := tc.KafkaClient.SendWithReply(lib.TestReplyTopic, message)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, string(response))
}

func (tc *TaskController) TestConsumeTopic(replyTopic string, message []byte) {
	fmt.Println("getting message from test topic:", string(message))
	tc.KafkaClient.Send(replyTopic, []byte("all ok"))
}

func (tc *TaskController) TestSecondTopic(replyTopic string, message []byte) {
	fmt.Println("getting message from second topic:", string(message))
}
