package usecase

import (
	"context"

	"main/domain"
	"main/internal/tokenutil"
	"main/lib"
)

type TaskUsecase struct {
	taskInfrastructure domain.TaskInfrastructure
	env                lib.Env
}

func NewTaskUsecase(taskInfrastructure domain.TaskInfrastructure, env lib.Env) domain.TaskUsecase {
	return &TaskUsecase{
		taskInfrastructure: taskInfrastructure,
		env:                env,
	}
}

func (tu *TaskUsecase) Create(c context.Context, task *domain.Task) error {
	return tu.taskInfrastructure.Create(c, task)
}

func (tu *TaskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	return tu.taskInfrastructure.FetchByUserID(c, userID)
}

func (u *TaskUsecase) ExtractIDFromToken(requestToken string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, u.env.RefreshTokenSecret)
}
