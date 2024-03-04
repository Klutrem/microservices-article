package task_controller

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTaskController),
	fx.Provide(NewTaskRouter),
)
