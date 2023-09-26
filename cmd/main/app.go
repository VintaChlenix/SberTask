package main

import (
	"SberTask/configs"
	"SberTask/internal/app"
	"SberTask/internal/db/postgres"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func run(log *zap.SugaredLogger) error {
	config, err := configs.GetConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	dbClient, err := postgres.NewClient(config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Infoln("connected to database")

	sberApp, err := app.NewApp(log, dbClient, config)
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}
	log.Infoln("app initialized")
	log.Infof("Starting server on: %s", config.Server.URL)
	if err := http.ListenAndServe(config.Server.URL, newRouter(sberApp)); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := run(sugar); err != nil {
		sugar.Error(err.Error())
		os.Exit(1)
	}
}
