package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/app"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/config"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/delivery/http/router"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/infra/mongodb"
)

func init() {
	serverMode := os.Getenv("SERVER_MODE")

	if serverMode == "" {
		log.Fatalln("missing or empty SERVER_MODE env variable ,please sect it dev for developement mod or prod for the production mode")
	}

	if serverMode == "dev" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("failed to load the env varibles , Error : ", err.Error())
		}
		return
	}
	if serverMode != "prod" {
		log.Fatalln("invalid SERVER_MODE env variable ,set it to dev for  devlopemnet mode or prod for the production mode")
	}
}
func main() {
	config := config.InitConfig()
	mongoConn := mongodb.ConnectToMongoDB(config.MongodbURI)
	mongoConn.Checkconnection()
	defer mongoConn.CloseConnetion()

	e := echo.New()
	messageService := app.NewMessageService(mongoConn.Client)

	router.InitRoutes(e, messageService)

	e.Logger.Fatal(e.Start(config.ServerAdress))

}
