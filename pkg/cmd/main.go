package cmd

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"super_api/pkg/conf"
)

var (
	cfg     *conf.Config
	logger  *zap.Logger
	command = &cobra.Command{}
)

func Execute() {
	logger, _ = zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync() // flushes buffer, if any

	cfg = &conf.Config{}
	if err := envconfig.Process("SUPER_API", cfg); err != nil {
		logger.Fatal(fmt.Sprintf("Config init failed with error: %s\n", err))
	}

	if err := command.Execute(); err != nil {
		logger.Fatal("Command execution failed with error", zap.Error(err))
		os.Exit(1)
	}
}
