package infrastructure

import (
	"context"

	"main/domain"
	"main/lib"
	"main/pkg/postgresql"
)

type userInfrastructure struct {
	database postgresql.Database
	table    string
	broker   lib.Rabbitbroker
}

func NewUserInfrastructure(db lib.PostgresClient, broker lib.Rabbitbroker) domain.UserInfrastructure {
	return &userInfrastructure{
		database: db.Client.Database(),
		table:    domain.UserTable,
		broker:   broker,
	}
}

func (ur *userInfrastructure) Create(c context.Context, user *domain.User) error {
	query := "INSERT INTO " + ur.table + " (email, password, name) VALUES (:email, :password, :name)"
	_, err := ur.database.NamedExec(query, user)
	return err
}

func (ur *userInfrastructure) Fetch(c context.Context) ([]domain.User, error) {
	query := "SELECT * FROM " + ur.table
	var users []domain.User
	err := ur.database.Select(&users, query)

	return users, err
}

func (ur *userInfrastructure) GetByEmail(c context.Context, email string) (domain.User, error) {
	query := "SELECT * FROM " + ur.table + " WHERE email = $1"
	var user domain.User
	err := ur.database.Get(&user, query, email)

	return user, err
}

func (ur *userInfrastructure) GetByID(c context.Context, id string) (domain.User, error) {
	query := "SELECT * FROM " + ur.table + " WHERE id = $1"
	var user domain.User
	err := ur.database.Get(&user, query, id)

	return user, err
}
