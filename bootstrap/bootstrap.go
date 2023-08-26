package bootstrap

import (
	"context"
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/api/middlewares"
	"github.com/BetterToPractice/go-gin-setup/api/routes"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"go.uber.org/fx"
)

var Module = fx.Options(
	lib.Module,
	controllers.Module,
	routes.Module,
	middlewares.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler lib.HttpHandler, routes routes.Routes, middlewares middlewares.Middlewares) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				middlewares.Setup()
				routes.Setup()
				if err := handler.Engine.Run(":8080"); err != nil {
					fmt.Println("error when run service", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error { return nil },
	})
}
