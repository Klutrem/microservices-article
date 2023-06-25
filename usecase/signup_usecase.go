package usecase

import (
	"context"
	"time"

	"main/domain"
	"main/internal/tokenutil"
)

type signupUsecase struct {
	userInfrastructure domain.UserInfrastructure
	contextTimeout     time.Duration
}

func NewSignupUsecase(userInfrastructure domain.UserInfrastructure, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userInfrastructure: userInfrastructure,
		contextTimeout:     timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userInfrastructure.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userInfrastructure.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
