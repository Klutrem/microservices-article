package infrastructure

import (
	"context"

	"main/domain"
	"main/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskInfrastructure struct {
	database   mongo.Database
	collection string
}

func NewTaskInfrastructure(db mongo.Database, collection string) domain.TaskInfrastructure {
	return &taskInfrastructure{
		database:   db,
		collection: collection,
	}
}

func (tr *taskInfrastructure) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *taskInfrastructure) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []domain.Task{}, err
	}

	return tasks, err
}
