package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	MongoHost string
	MongoPort string
	Database  string
}

func LoadConfig() (c Config, err error) {
	err = godotenv.Load("./config/envs/dev.env")
	if err != nil {
		log.Fatalf("Some error occued .env Err: %s", err)
	}
	c = Config{Port: os.Getenv("PORT"), MongoHost: os.Getenv("MONGO_HOST"), MongoPort: os.Getenv("MONGO_PORT"), Database: os.Getenv("MONGO_DATABASE")}
	return
}
