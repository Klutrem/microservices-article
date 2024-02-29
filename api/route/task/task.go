package TaskRoute

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTaskRouter),
	fx.Provide(NewKafkaHandler),
)
