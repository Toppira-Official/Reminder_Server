package main

import (
	"github.com/Toppira-Official/backend/internal/configs"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			fx.Provide(
				configs.GetEnvironments,
				configs.NewHttpServer,
			),
			fx.Invoke(
				configs.LoadEnvironmentsFromEnvFile,
			),
		).
		Run()
}
