package main

import (
	"log"
	"servicetemplate/internal/server"
	"servicetemplate/pkg/db"
	"servicetemplate/pkg/env"
	"servicetemplate/pkg/logger"
)

// @title Service Template API
// @version 1.0
// @description This is a sample server for a Service Template API.
// @termsOfService http://swagger.io/terms/

// @contact.name DevVictor19
// @contact.url https://www.linkedin.com/in/antonio-victor-borges-4a2852228/
// @contact.email antoniovictor12@live.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
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
