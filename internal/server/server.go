package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"servicetemplate/pkg/env"
	"servicetemplate/pkg/logger"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	cfg    *env.Config
	logger logger.Logger
	db     *sqlx.DB
}

func NewServer(cfg *env.Config, logger logger.Logger, db *sqlx.DB) *Server {
	return &Server{
		echo:   echo.New(),
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Addr,
		ReadTimeout:    s.cfg.Server.ReadTimeout,
		WriteTimeout:   s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: s.cfg.Server.MaxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Addr)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), s.cfg.Server.CtxTimeout)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
