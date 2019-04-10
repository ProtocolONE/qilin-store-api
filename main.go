package main

import (
	"fmt"
	"go.uber.org/zap"
	"super_api/pkg/conf"
	"super_api/pkg/server"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	config := &conf.Config{}

	if err := envconfig.Process("SUPER_API", config); err != nil {
		logger.Fatal(fmt.Sprintf("Config init failed with error: %s\n", err))
	}

	apiServer, err := server.NewServer(config)
	if err != nil {
		logger.Fatal("Failed to create apiServer", zap.Error(err))
	}
	err = apiServer.Start()
	if err != nil {
		logger.Fatal("Failed to start apiServer", zap.Error(err))
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Sugar().Error("Failed on main recover", "error", err)
		}
	}()
}