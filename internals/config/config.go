package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAdress string
	MongodbURI   string
	DatabaseName string
}

func InitConfig() *Config {
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		log.Fatalln("missing or empty SERVER_ADDRESS env variable ")
	}

	mongodbUrl := os.Getenv("MONGODB_URL")

	if mongodbUrl == "" {
		log.Fatalln("missing or empty MONGODB_URL env variable")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatalln("missing or empty DATABASE_NAME env variable")
	}
	return &Config{
		ServerAdress: serverAddress,
		MongodbURI:   mongodbUrl,
		DatabaseName: databaseName,
	}

}
