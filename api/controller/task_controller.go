package controller

import (
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

// I don't care, it won't have business logic, just consume the topic
func (tc *TaskController) TestConsumeTopic(topic string) {
	messages := tc.KafkaClient.Consume(topic)
	for message := range messages {
		fmt.Println(string(message.Value)) //replace this to service calling
	}
}
