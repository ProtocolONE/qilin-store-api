package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"super_api/pkg/server"
	"super_api/pkg/services"
)

func init() {
	serverCmd = &cobra.Command{
		Use:   "event-bus",
		Short: "Run event bus listener",
		Run:   runEventBus,
	}
	command.AddCommand(serverCmd)
}

func runEventBus(_ *cobra.Command, _ []string) {
	zap.L().Info("Starting event bus command")

	dbProvider, err := services.NewDatabaseProvider(cfg.Db.Connection, cfg.Db.Name)
	if err != nil {
		zap.L().Fatal("Can't create database", zap.Error(err))
	}

	bus, err := server.NewEventBus(dbProvider, cfg.Bus.Connection)
	if err != nil {
		zap.L().Fatal("Can't create event bus", zap.Error(err))
	}
	defer bus.Shutdown()

	zap.L().Info(fmt.Sprintf("Starting up event bus worker. Connection: %s", cfg.Bus.Connection))
	if err = bus.StartListen(); err != nil {
		zap.L().Fatal("Error running event bus worker", zap.Error(err))
	}
}