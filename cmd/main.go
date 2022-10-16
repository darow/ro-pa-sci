package main

import (
	"github.com/darow/ro-pa-sci/internal/server"
	"log"

	"go.uber.org/zap"
)

// @title ro-pa-sci API
// @version 1.0
// @description API server for rock-paper-scissors game

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name session

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := server.NewConfig("config.yml")
	if err != nil {
		logger.Sugar().Error(err)
	}

	log.Fatalln(server.Start(cfg, logger.Sugar()))
}
