package usecase

import (
	"context"
	"time"

	"main/domain"
	"main/internal/tokenutil"
)

type loginUsecase struct {
	userInfrastructure domain.UserInfrastructure
	contextTimeout     time.Duration
}

func NewLoginUsecase(userInfrastructure domain.UserInfrastructure, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userInfrastructure: userInfrastructure,
		contextTimeout:     timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userInfrastructure.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
