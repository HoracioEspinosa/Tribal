package helpers

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PublicConfig = Config{}

type Config struct {
	Host string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User string `env:"POSTGRES_USER" envDefault:"tribal"`
	Pwd  string `env:"POSTGRES_PASSWORD" envDefault:"secret"`
	Name string `env:"POSTGRES_DB" envDefault:"tribal"`
}

func Connect(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.User, cfg.Pwd, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	return db
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg
}
