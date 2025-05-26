package config

import (
	"os"
	"slices"

	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type Config struct {
	MongoDb struct {
		Host     string `env:"MONGO_HOST,required" usage:"mongo database host"`
		Username string `env:"MONGO_USERNAME,required" usage:"mongo database username"`
		Password string `env:"MONGO_PASSWORD,required" usage:"mongo database password"`
		Database string `env:"MONGO_DATABASE,required" usage:"mongo database name"`
		Port     int    `env:"MONGO_PORT,required" usage:"mongo database port"`
		Uri      string `env:"MONGO_URI,required,expand" usage:"mongo database uri"`
	} `env:"DB"`
	Http struct {
		Port int `env:"PORT,required" usage:"http service port"`
	} `env:"HTTP"`
	Jwt struct {
		SecretKey string `env:"SECRET_KEY,required" usage:"jwt secret key"`
		ExpireSec int    `env:"EXPIRE_SEC,required" usage:"jwt expire sec"`
	} `env:"JWT"`
}

func NewConfig() (cfg *Config, err error) {
	runtimeEnv := os.Getenv("GO_ENV")
	if !slices.Contains([]string{"production", "prod"}, runtimeEnv) {
		if err = godotenv.Load(); err != nil {
			return nil, err
		}
	}

	cfg = &Config{}
	if err = env.Load(cfg, &env.Options{NameSep: "_"}); err != nil {
		return nil, err
	}

	return
}
