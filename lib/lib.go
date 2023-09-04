package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewHttpHandler),
	fx.Provide(NewSwagger),
	fx.Provide(NewDatabase),
	fx.Provide(NewMigration),
)
