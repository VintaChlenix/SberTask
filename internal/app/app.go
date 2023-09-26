package app

import (
	"SberTask/configs"
	"SberTask/internal/db"
	"go.uber.org/zap"
)

type App struct {
	log    *zap.SugaredLogger
	db     db.DB
	config *configs.Config
}

func NewApp(log *zap.SugaredLogger, db db.DB, config *configs.Config) (*App, error) {
	return &App{
		log:    log,
		db:     db,
		config: config,
	}, nil
}
