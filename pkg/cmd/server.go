package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"super_api/pkg/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Qilin-Store api server with given configuration",
	Run:   runServer,
}

func init() {
	command.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) {
	apiServer, err := server.NewServer(cfg)
	if err != nil {
		zap.L().Fatal("Failed to create server", zap.Error(err))
	}

	defer apiServer.Shutdown()

	zap.L().Info("Starting up server")
	if err = apiServer.Start(); err != nil {
		zap.L().Fatal("Error running server", zap.Error(err))
	}
}

