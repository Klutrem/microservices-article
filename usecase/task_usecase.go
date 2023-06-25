package usecase

import (
	"context"
	"time"

	"main/domain"
)

type taskUsecase struct {
	taskInfrastructure domain.TaskInfrastructure
	contextTimeout     time.Duration
}

func NewTaskUsecase(taskInfrastructure domain.TaskInfrastructure, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskInfrastructure: taskInfrastructure,
		contextTimeout:     timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskInfrastructure.Create(ctx, task)
}

func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskInfrastructure.FetchByUserID(ctx, userID)
}
