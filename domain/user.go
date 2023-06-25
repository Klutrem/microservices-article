package domain

import (
	"context"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserInfrastructure interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
