package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"main/domain"
	"main/lib"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
	KafkaClient lib.KafkaClient
}

func NewTaskController(usecase domain.TaskUsecase, kafka lib.KafkaClient) TaskController {
	return TaskController{
		TaskUsecase: usecase,
		KafkaClient: kafka,
	}
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task

	authHeader := c.Request.Header.Get("Authorization")
	t := strings.Split(authHeader, " ")
	userID, err := tc.TaskUsecase.ExtractIDFromToken(t[1])
	if err != nil {
		log.Println(userID)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	task.UserID, err = strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (tc *TaskController) Fetch(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	t := strings.Split(authHeader, " ")
	userID, err := tc.TaskUsecase.ExtractIDFromToken(t[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	tasks, err := tc.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) TestReplyTopic(c *gin.Context) {
	message, err := json.Marshal("test_message")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	response, err := tc.KafkaClient.SendWithReply(lib.TestReplyTopic, message) // shouldn't be like that, it should call service, and service should send to kafka
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, string(response)) //it's better to unmarshal response to struct (json.Unmarshal(response, &struct))
}

// I don't care, it won't have business logic, just handling topic, but for real functions it must call service
// Don't forget to send reply on the reply topic
func (tc *TaskController) TestConsumeTopic(replyTopic string, message []byte) {
	fmt.Println("getting message from test topic:", string(message))
	tc.KafkaClient.Send(replyTopic, []byte("all ok"))
}

func (tc *TaskController) TestSecondTopic(replyTopic string, message []byte) {
	fmt.Println("getting message from second topic:", string(message))
}
