package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type dbcfg struct {
	DSN         string `mapstructure:"USER_SERVICE_DB_DSN"`
	DBName      string `mapstructure:"USER_SERVICE_DB_NAME"`
	MinPoolSize int    `mapstructure:"USER_SERVICE_DB_MIN_POOL_SIZE"`
	MaxPoolSize int    `mapstructure:"USER_SERVICE_DB_MAX_POOL_SIZE"`
	MaxIdleTime string `mapstructure:"USER_SERVICE_DB_MAX_IDLE_TIME"`
}

type Config struct {
	Port        int    `mapstructure:"USER_SERVICE_PORT"`
	Environment string `mapstructure:"USER_SERVICE_ENVIRONMENT"`
	Db          dbcfg  `mapstructure:",squash"`
}

func New() (*Config, error) {
	cfg := &Config{}

	viper.SetEnvPrefix("USER_SERVICE")

	viper.SetDefault("USER_SERVICE_PORT", os.Getenv("USER_SERVICE_PORT"))
	viper.SetDefault("USER_SERVICE_ENVIRONMENT", "development")

	viper.SetDefault("USER_SERVICE_DB_MIN_POOL_SIZE", 25)
	viper.SetDefault("USER_SERVICE_DB_MAX_POOL_SIZE", 25)
	viper.SetDefault("USER_SERVICE_DB_MAX_IDLE_TIME", "15m")

	viper.SetDefault("USER_SERVICE_DB_DSN", os.Getenv("USER_SERVICE_DB_DSN"))
	viper.SetDefault("USER_SERVICE_DB_NAME", os.Getenv("USER_SERVICE_DB_NAME"))

	viper.AutomaticEnv()

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	if err := checkUnset(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.IgnoreUntaggedFields = true
		dc.ErrorUnset = true
	}); err != nil {
		return nil, err
	}

	if structs.HasZero(&cfg) {
		return nil, fmt.Errorf("config type has zero value")
	}
	return cfg, nil
}

func checkUnset() error {
	var listunset = []string{}
	for key, val := range viper.AllSettings() {
		if val == "" {
			listunset = append(listunset, strings.ToUpper(key))
		}
	}

	if len(listunset) != 0 {
		envs := pie.Join(listunset, " | ")
		return fmt.Errorf("ENVIRONMENT NOT SET: %v", envs)
	}
	return nil
}
