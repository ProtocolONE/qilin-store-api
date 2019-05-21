package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"github.com/ProtocolONE/qilin-store-api/pkg/server"
	"github.com/ProtocolONE/qilin-store-api/pkg/services"
)

func init() {
	eventBusCmd := &cobra.Command{
		Use:   "event-bus",
		Short: "Run event bus listener",
		Run:   runEventBus,
	}
	command.AddCommand(eventBusCmd)
}

func runEventBus(_ *cobra.Command, _ []string) {
	zap.L().Info("Starting event bus command")

	dbProvider, err := services.NewDatabaseProvider(cfg.Db)
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
