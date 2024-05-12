package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/api"
	dbuser "github.com/perfectgentlemande/go-openapi-generator-example/internal/database"
)

type Config struct {
	Server *api.Config
	DBUser *dbuser.Config
}

func readConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("cannot load .env file: %w", err)
	}

	return Config{
		Server: &api.Config{
			Addr: os.Getenv("SERVER_ADDR"),
		},
		DBUser: &dbuser.Config{
			DBName:  os.Getenv("DBUSER_DBNAME"),
			ConnStr: os.Getenv("DBUSER_CONNSTR"),
		},
	}, nil
}
