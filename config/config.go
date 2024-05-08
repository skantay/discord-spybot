package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Discord Discord `yaml:"discord"`
	S3      S3      `yaml:"s3"`
}

type Discord struct {
	Token string `yaml:"token"`
}

type S3 struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Region    string `yaml:"region"`
}

func Get(path string) (Config, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
