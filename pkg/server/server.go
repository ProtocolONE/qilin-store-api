package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"strconv"
	"github.com/ProtocolONE/qilin-store-api/pkg/api"
	"github.com/ProtocolONE/qilin-store-api/pkg/conf"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/services"
)

type server struct {
	echo         *echo.Echo
	serverConfig *conf.ServerConfig
	db           interfaces.DatabaseProvider
}

func NewServer(config *conf.Config) (*server, error) {
	httpServer := echo.New()
	server := &server{
		echo:         httpServer,
		serverConfig: config.Server,
	}

	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())

	httpServer.HTTPErrorHandler = server.QilinStoreErrorHandler

	dbProvider, err := services.NewDatabaseProvider(config.Db)
	if err != nil {
		return nil, err
	}

	gameService := services.NewGameService(dbProvider)

	apiGroup := httpServer.Group("/api/v1")
	if _, err := api.InitGameRouter(apiGroup, gameService); err != nil {
		return nil, err
	}

	return server, nil
}

func (s *server) Start() error {
	zap.L().Info("Starting http server", zap.Int("port", s.serverConfig.Port))

	return s.echo.Start(":" + strconv.Itoa(s.serverConfig.Port))
}

func (s *server) Shutdown() {
	s.db.Shutdown()
}
