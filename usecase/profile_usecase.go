package usecase

import (
	"context"
	"time"

	"main/domain"
)

type profileUsecase struct {
	userInfrastructure domain.UserInfrastructure
	contextTimeout     time.Duration
}

func NewProfileUsecase(userInfrastructure domain.UserInfrastructure, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userInfrastructure: userInfrastructure,
		contextTimeout:     timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userInfrastructure.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Email: user.Email}, nil
}
