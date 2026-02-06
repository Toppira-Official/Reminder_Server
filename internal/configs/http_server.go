package configs

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type HttpServer struct {
	envs Environments
}

func NewHttpServer(lc fx.Lifecycle, envs Environments) *gin.Engine {
	engine := gin.Default()

	lc.Append(
		fx.StartHook(func(ctx context.Context) {
			engine.Run(httpServerPortNumber(envs.PORT.String()))
		}),
	)

	return engine
}

func httpServerPortNumber(port string) string {
	return fmt.Sprintf(":%s", port)
}
