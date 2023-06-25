package domain

import (
	"context"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID     int    `bson:"_id" json:"-"`
	Title  string `bson:"title" form:"title" binding:"required" json:"title"`
	UserID int    `bson:"userID" json:"-"`
}

type TaskInfrastructure interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
