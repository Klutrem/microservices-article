package domain

import (
	"context"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	GetUserByID(c context.Context, email string) (User, error)
	CreateAccessToken(user *User) (accessToken string, err error)
	CreateRefreshToken(user *User) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string) (string, error)
}
