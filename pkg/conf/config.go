package conf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Config the application's configuration
type Config struct {
	Auth1  *Auth1
	Server *ServerConfig
	Db     *DbConfig
	Bus    *EventBusConfig
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("QILINSTOREAPI", cfg); err != nil {
	 	zap.L().Error(fmt.Sprintf("Config init failed with error: %s\n", err))
	 	return nil, err
	}
	return cfg, nil
}
