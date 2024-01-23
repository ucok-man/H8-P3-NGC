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

type jwtcfg struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

type Config struct {
	Port            int    `mapstructure:"GATEWAY_PORT"`
	UserServicePort int    `mapstructure:"USER_SERVICE_PORT"`
	Environment     string `mapstructure:"GATEWAY_ENVIRONMENT"`
	Jwt             jwtcfg `mapstructure:",squash"`
}

func New() (*Config, error) {
	cfg := &Config{}

	// viper.SetEnvPrefix("GATEWAY")

	viper.SetDefault("GATEWAY_PORT", os.Getenv("GATEWAY_PORT"))
	viper.SetDefault("GATEWAY_ENVIRONMENT", "development")
	viper.SetDefault("USER_SERVICE_PORT", os.Getenv("USER_SERVICE_PORT"))

	viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))

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
