package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

// A Conf struct wuth embedded serverConf to get the
// initial server configs.
type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
	Db     dbConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type dbConf struct {
	Host     string `env:"DB_HOST,requred"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
	DbName   string `env:"DB_NAME,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
