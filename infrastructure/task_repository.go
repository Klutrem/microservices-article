package infrastructure

import (
	"context"

	"main/domain"
	"main/pkg/postgresql"
)

type taskInfrastructure struct {
	db    postgresql.Database
	table string
}

func NewTaskInfrastructure(db postgresql.Database, table string) domain.TaskInfrastructure {
	return &taskInfrastructure{
		db:    db,
		table: table,
	}
}

func (tr *taskInfrastructure) Create(c context.Context, task *domain.Task) error {
	query := "INSERT INTO " + tr.table + " (user_id, name, description) VALUES (:user_id, :name, :description)"
	_, err := tr.db.NamedExec(query, task)

	return err
}

func (tr *taskInfrastructure) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	query := "SELECT * FROM " + tr.table + " WHERE user_id = $1"
	var tasks []domain.Task
	err := tr.db.Select(&tasks, query, userID)

	return tasks, err
}
