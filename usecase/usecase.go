package usecase

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserUsecase),
	fx.Provide(NewTaskUsecase),
)
