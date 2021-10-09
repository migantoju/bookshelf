package main

import (
	"bookshelf/config"
	"bookshelf/server/router"
	"fmt"
	"net/http"

	lr "bookshelf/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting Server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server Startup Failed...")
	}
}
