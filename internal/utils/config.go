package utils

import (
	"errors"
	"os"

	"github.com/mbilarusdev/fool_base/v2/log"
)

var Conf Config

type Config struct {
	LogLevel string
	PGX      string
	RDB      string
	Secret   string
}

func ParseConfig() *Config {
	conf := new(Config)
	conf.PGX = os.Getenv("PGX")
	if conf.PGX == "" {
		err := errors.New("ENV Var PGX doesn't exist or incorrect")
		log.Err(err, "")
		panic(err)
	}
	conf.RDB = os.Getenv("RDB")
	if conf.RDB == "" {
		err := errors.New("ENV Var RDB doesn't exist or incorrect")
		log.Err(err, "")
		panic(err)
	}
	conf.Secret = os.Getenv("SECRET")
	if conf.Secret == "" || len(conf.Secret) != 32 {
		err := errors.New("ENV Var secret doesn't exist or incorrect")
		log.Err(err, "")
		panic(err)
	}
	conf.LogLevel = os.Getenv("LOG_LEVEL")
	if conf.LogLevel == "" {
		conf.LogLevel = "debug"
	}
	return conf
}
