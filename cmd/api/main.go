package main

import (
	"log"
	"servicetemplate/internal/server"
	"servicetemplate/pkg/db"
	"servicetemplate/pkg/env"
	"servicetemplate/pkg/logger"
)

func main() {
	log.Println("Starting api server...")

	cfg := env.NewConfig()

	logger := logger.NewZapLogger(cfg)
	logger.InitLogger()
	logger.Infof("LogLevel: %s, Mode: %s",
		cfg.Logger.Level,
		cfg.Server.Mode,
	)

	mysqlDB := db.NewMySqlDB(cfg)
	logger.Infof("MySQL connected, Status: %#v", mysqlDB.Stats())
	defer mysqlDB.Close()

	s := server.NewServer(cfg, logger)
	if err := s.Start(); err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
