package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"main/domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func NewTaskController(usecase domain.TaskUsecase) TaskController {
	return TaskController{
		TaskUsecase: usecase,
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
