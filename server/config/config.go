package config

import "github.com/spf13/viper"

type Config struct {
	DB Database
}

type Database struct {
	DataSourceName     string
	MaxConnections     int
	MaxIdleConnections int
}

func Import() (Config, error) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
