package main

import (
	"ewallet-backend/internal/config"
	"ewallet-backend/pkg/logger"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	log, err := logger.New(cfg.Log.Level, cfg.Log.Format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logger: %v\n", err)
		os.Exit(1)
	}

	defer log.Sync()

	log.Info("Starting application",
		zap.String("name", cfg.App.Name),
		zap.String("env", cfg.App.Env),
	)
}
