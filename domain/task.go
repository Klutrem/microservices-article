package domain

import (
	"context"
)

const (
	TaskTable = "tasks"
)

type Task struct {
	ID     int    `db:"id" json:"-"`
	Title  string `db:"title" form:"title" binding:"required" json:"title"`
	UserID int    `db:"userid" json:"-"`
}

type TaskInfrastructure interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
	ExtractIDFromToken(requestToken string) (string, error)
}
