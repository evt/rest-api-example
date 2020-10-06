package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is a config :)
type Config struct {
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"PG_URL"`
	PgProto          string `envconfig:"PG_PROTO"`
	PgAddr           string `envconfig:"PG_ADDR"`
	PgDb             string `envconfig:"PG_DB"`
	PgUser           string `envconfig:"PG_USER"`
	PgPassword       string `envconfig:"PG_PASSWORD"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	PgCertPath       string `envconfig:"PG_CERT_PATH"`
	PgLocation       string `envconfig:"PG_LOCATION"`
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	GCBucket         string `envconfig:"GC_BUCKET"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
