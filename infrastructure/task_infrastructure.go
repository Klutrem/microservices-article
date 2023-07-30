package infrastructure

import (
	"context"
	"strconv"

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
	query := "INSERT INTO " + tr.table + " (userid, title) VALUES (:userid, :title)"
	_, err := tr.db.NamedExec(query, task)

	return err
}

func (tr *taskInfrastructure) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	query := "SELECT * FROM " + tr.table + " WHERE userid = ($1)"
	var tasks []domain.Task
	intid, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	err = tr.db.Select(&tasks, query, intid)

	return tasks, err
}
