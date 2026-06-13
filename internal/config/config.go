package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP HTTPConfig `yaml:"http"`
	DB   DBConfig   `yaml:"database"`
}

type HTTPConfig struct {
	Port string `yaml:"port" env-default:"8080"`
}

type DBConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

func Load() (Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return Config{}, errors.New("Не указан путь до конфиг файла")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("Ошибка файл не существует, или путь до файла указан не правильно: %w", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return Config{}, fmt.Errorf("Ошибка при прочтении конфига: %w", err)
	}

	return cfg, nil
}
