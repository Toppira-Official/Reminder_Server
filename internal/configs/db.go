package configs

import (
	"context"
	"time"

	"github.com/glebarez/sqlite"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(lc fx.Lifecycle, envs Environments, log *zap.Logger) *gorm.DB {
	var sqliteFileName string

	switch envs.MODE.String() {
	case "production":
		sqliteFileName = "production.db"
	default:
		sqliteFileName = "dev.db"
	}

	gormLogger := logger.New(
		zap.NewStdLog(log),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(sqlite.Open(sqliteFileName),
		&gorm.Config{
			Logger:                                   gormLogger,
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		log.Fatal("failed to connect to db", zap.Error(err))
	}

	db.Exec("PRAGMA journal_mode = WAL;")
	db.Exec("PRAGMA foreign_keys = ON;")
	db.Exec("PRAGMA busy_timeout = 5000;")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql db", zap.Error(err))
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	lc.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				log.Info("closing database connection")
				return sqlDB.Close()
			},
		},
	)

	return db
}
