package configs

import (
	"context"
	"os"
	"path/filepath"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(lc fx.Lifecycle, envs Environments) *zap.Logger {
	var logger *zap.Logger

	switch envs.MODE.String() {
	case "production":
		config := zap.NewProductionConfig()
		config.DisableStacktrace = true
		config.OutputPaths = []string{"stdout", envs.LOG_FILE.String()}
		config.ErrorOutputPaths = []string{"stderr", envs.LOG_FILE.String()}
		_ = os.MkdirAll(filepath.Dir(envs.LOG_FILE.String()), 0755)
		logger, _ = config.Build()
	default:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.OutputPaths = []string{"stderr"}
		logger, _ = config.Build()
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger
}
