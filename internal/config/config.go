package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ConfigPathVar = "CONFIG_PATH"
	JWTSecretVar  = "JWT_SECRET"
	DBUserVar     = "DB_USER"
	DBPasswordVar = "DB_PASSWORD"
	DBHostVar     = "DB_HOST"
	DBPortVar     = "DB_PORT"
	DBDatabaseVar = "DB_DATABASE"
	EnvLocal      = "local"
	EnvDev        = "dev"
	EnvProd       = "prod"
)

type Config struct {
	Env         string        `yaml:"env" env-required:"true"`
	JWTDuration time.Duration `yaml:"jwt_duration"`
	JWTSecret   string        `yaml:"-"`
	HTTPServer  `yaml:"http_server"`
	DBParam     `yaml:"-"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
	Prefix      string        `yaml:"prefix" env-default:"/"`
}

type DBParam struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

var Cfg *Config

func MustLoad() *Config {
	configPath, isFound := os.LookupEnv(ConfigPathVar)
	if !isFound {
		log.Fatalf("The environment variable %s is not set", ConfigPathVar)
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Cannot find config file %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err.Error())
	}

	cfg.JWTSecret, isFound = os.LookupEnv(JWTSecretVar)
	if !isFound {
		log.Fatalf(".env file must contain %s variable", JWTSecretVar)
	}

	cfg.DBParam = DBParam{
		User:     os.Getenv(DBUserVar),
		Password: os.Getenv(DBPasswordVar),
		Host:     os.Getenv(DBHostVar),
		Port:     os.Getenv(DBPortVar),
		DB:       os.Getenv(DBDatabaseVar),
	}

	return &cfg
}
