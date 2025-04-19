package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"Http_Server"`
}

func MustLoad () *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if (configPath == "") {
		flags := flag.String("config", "", "")
		flag.Parse()

		configPath = *flags

		if (configPath == "") {
			log.Fatal("config path doesn't set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("cofig file does not exist")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if (err != nil) {
		log.Fatal("config error", err.Error())
	}

	return &cfg
}