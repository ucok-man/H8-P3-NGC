package config

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func init() {
	godotenv.Load()
}

type dbcfg struct {
	DSN         string `mapstructure:"DB_DSN"`
	MaxOpenConn int    `mapstructure:"DB_MAX_OPEN_CONN"`
	MaxIdleConn int    `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxIdleTime string `mapstructure:"DB_MAX_IDLE_TIME"`
}

type jwtcfg struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

type Config struct {
	Port        int    `mapstructure:"API_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Db          dbcfg  `mapstructure:",squash"`
	Jwt         jwtcfg `mapstructure:",squash"`
}

func New() (*Config, error) {
	cfg := &Config{}

	viper.SetDefault("API_PORT", 8080)
	viper.SetDefault("ENVIRONMENT", "development")

	// viper.SetDefault("DB_DSN", os.Getenv("DB_DSN"))
	viper.SetDefault("DB_MAX_OPEN_CONN", 25)
	viper.SetDefault("DB_MAX_IDLE_CONN", 25)
	viper.SetDefault("DB_MAX_IDLE_TIME", "15m")

	// viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.IgnoreUntaggedFields = true
		dc.ErrorUnused = true
		dc.ErrorUnset = true
	}); err != nil {
		return nil, err
	}

	if structs.HasZero(&cfg) {
		return nil, fmt.Errorf("config type has zero value")
	}
	return cfg, nil
}
