package server

import (
	"fmt"
	"github.com/ProtocolONE/authone-jwt-verifier-golang"
	redisStorage "github.com/ProtocolONE/authone-jwt-verifier-golang/storage/redis"
	"github.com/ProtocolONE/qilin-store-api/pkg/api"
	"github.com/ProtocolONE/qilin-store-api/pkg/conf"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/ProtocolONE/qilin-store-api/pkg/services"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"strconv"
)

type server struct {
	echo         *echo.Echo
	serverConfig *conf.ServerConfig
	db           interfaces.DatabaseProvider
}

type StoreValidator struct {
	validator *validator.Validate
}

func (cv *StoreValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewServer(config *conf.Config) (*server, error) {
	httpServer := echo.New()
	server := &server{
		echo:         httpServer,
		serverConfig: config.Server,
	}

	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(session.Middleware(sessions.NewCookieStore([]byte(config.Sessions.Secret))))

	settings := jwtverifier.Config{
		ClientID:     config.Auth1.ClientId,
		ClientSecret: config.Auth1.ClientSecret,
		Scopes:       []string{"openid", "offline"},
		RedirectURL:  "",
		Issuer:       config.Auth1.Issuer,
	}

	httpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders:    []string{"x-centrifugo-token", "x-items-count"},
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowOrigins:     config.Server.AllowOrigins,
		AllowCredentials: config.Server.AllowCredentials,
	}))
	httpServer.Pre(middleware.RemoveTrailingSlash())

	httpServer.HTTPErrorHandler = server.QilinStoreErrorHandler

	validate := validator.New()
	httpServer.Validator = &StoreValidator{validator: validate}

	dbProvider, err := services.NewDatabaseProvider(config.Db)
	if err != nil {
		return nil, err
	}

	gameService := services.NewGameService(dbProvider)

	apiGroup := httpServer.Group("/api/v1")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Sessions.Host, config.Sessions.Port),
		Password: config.Sessions.Password,
		DB:       0, // use default DB
	})
	err = client.Ping().Err()
	if err != nil {
		return nil, err
	}

	adapter, err := redisStorage.NewStorage(client)
	if err != nil {
		return nil, err
	}
	jwtv := jwtverifier.NewJwtVerifier(settings)
	jwtv.SetStorage(adapter)
	apiGroup.Use(AuthOneJwtWithConfig(jwtv))

	if _, err := api.InitGameRouter(apiGroup, gameService); err != nil {
		return nil, err
	}

	accountService := services.NewAccountService(dbProvider)
	if _, err := api.InitAccountRouter(apiGroup, accountService); err != nil {
		return nil, err
	}

	profileService := services.NewProfileService(dbProvider)
	if _, err := api.InitProfileRouter(apiGroup, profileService); err != nil {
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
