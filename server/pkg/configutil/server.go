package configutil

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

type ServerConfig struct {
	Host    string `envconfig:"HOST" default:"0.0.0.0"`
	Port    string `envconfig:"PORT" default:"18080"`
	BaseURL string `envconfig:"BASE_URL"`
	Deploy  string `envconfig:"DEPLOY"`
	MySQL   MySQLConfig
	//Redis   RedisConfig
	//Session SessionConfig
	//AWS     AWSConfig
}

var (
	sc ServerConfig
)

func ServerInit() error {
	err := envconfig.Process("", &sc)
	if err != nil {
		return xerrors.Errorf("message: %w", err)
	}
	return nil
}

func GetServerConfig() ServerConfig {
	return sc
}
