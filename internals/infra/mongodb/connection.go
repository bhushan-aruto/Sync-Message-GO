package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Client *mongo.Client
}

func ConnectToMongoDB(mongodbURI string) *MongoConnection {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))

	if err != nil {
		log.Fatalln("Error occurred while connecting to MongoDB, Error:", err.Error())
	}

	return &MongoConnection{
		Client: client,
	}
}

func (m *MongoConnection) Checkconnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Ping(ctx, nil); err != nil {
		log.Println("error occured with the mongodb , Error :", err.Error())
		return
	}

	log.Println("connected to mongodb")

}

func (m *MongoConnection) CloseConnetion() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		log.Fatalln("error while closing the mongodb connection ,Error : ", err.Error())
	}

	log.Println(" mongobd connection closed succesfully")
}
