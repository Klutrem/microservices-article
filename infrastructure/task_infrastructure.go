package infrastructure

import (
	"context"

	"main/domain"
	"main/lib"
	"main/pkg/postgresql"
)

type taskInfrastructure struct {
	db    postgresql.Database
	table string
}

func NewTaskInfrastructure(db lib.PostgresClient) domain.TaskInfrastructure {
	return &taskInfrastructure{
		db:    db.Client.Database(),
		table: domain.TaskTable,
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
