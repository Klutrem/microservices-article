package usecase

import (
	"context"

	"main/domain"
	"main/internal/tokenutil"
	"main/lib"
)

type UserUsecase struct {
	userInfrastructure domain.UserInfrastructure
	env                lib.Env
}

func NewUserUsecase(userInfrastructure domain.UserInfrastructure, env lib.Env) domain.UserUsecase {
	return UserUsecase{
		userInfrastructure: userInfrastructure,
		env:                env,
	}
}

func (u UserUsecase) Create(c context.Context, user *domain.User) error {
	return u.userInfrastructure.Create(c, user)
}

func (u UserUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	return u.userInfrastructure.GetByEmail(c, email)
}

func (u UserUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {

	user, err := u.userInfrastructure.GetByID(c, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Email: user.Email}, nil
}

func (u UserUsecase) GetUserByID(c context.Context, email string) (domain.User, error) {
	return u.userInfrastructure.GetByID(c, email)
}

func (u UserUsecase) CreateAccessToken(user *domain.User) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, u.env.AccessTokenSecret, u.env.AccessTokenExpiryHour)
}

func (u UserUsecase) CreateRefreshToken(user *domain.User) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, u.env.RefreshTokenSecret, u.env.RefreshTokenExpiryHour)
}

func (u UserUsecase) ExtractIDFromToken(requestToken string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, u.env.RefreshTokenSecret)
}
