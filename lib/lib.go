package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewHttpHandler),
	fx.Provide(NewMail),
	fx.Provide(NewDatabase),
	fx.Provide(NewMigration),
)
