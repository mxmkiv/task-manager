package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	ServicePort     string
	SecretKey       string
	DbConnectionUrl string
}

// db link config

func NewServiceConfig() *ServiceConfig {

	godotenv.Load(".env")

	key, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Fatal("failed to load key from env")
	}

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = ":8888"
		log.Println("failed to load port from env, service start on default :8888")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("failed to load db user from env")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Fatal("failed to load db password from env")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("failed to load db host from env")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		log.Fatal("failed to load db port from env")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("failed to load db name from env")
	}

	return &ServiceConfig{
		ServicePort:     port,
		SecretKey:       key,
		DbConnectionUrl: "postgresql://" + user + ":" + password + "@" + host + ":" + dbPort + "/" + dbName,
	}

}
