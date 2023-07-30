package lib

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPostgresClient),
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewKafkaClient),
)
