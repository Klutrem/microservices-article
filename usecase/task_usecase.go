package usecase

import (
	"context"

	"main/domain"
)

type TaskUsecase struct {
	taskInfrastructure domain.TaskInfrastructure
}

func NewTaskUsecase(taskInfrastructure domain.TaskInfrastructure) domain.TaskUsecase {
	return &TaskUsecase{
		taskInfrastructure: taskInfrastructure,
	}
}

func (tu *TaskUsecase) Create(c context.Context, task *domain.Task) error {
	return tu.taskInfrastructure.Create(c, task)
}

func (tu *TaskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	return tu.taskInfrastructure.FetchByUserID(c, userID)
}
