package lib

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewKafkaClient),
	fx.Provide(GetLogger),
)
