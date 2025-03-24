package app

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vithsutra/ca-chat-sync-message-service/internals/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageService struct {
	DB *mongo.Client
}

func NewMessageService(db *mongo.Client) *MessageService {
	return &MessageService{
		DB: db,
	}
}

func (s *MessageService) GetOrderedMessages(userId string, currentMessageId string) ([]domain.Message, error) {
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatalln("empty or missing DATABASE_NAME env variable")
	}

	collection := s.DB.Database(databaseName).Collection(userId)
	// log.Println("Fetching messages from Collection:", userId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentID, err := strconv.Atoi(currentMessageId)
	if err != nil {
		// log.Println("Invalid currentMessageId:", err)
		return nil, err
	}

	filter := bson.M{"message_id": bson.M{"$gt": strconv.Itoa(currentID)}}
	opts := options.Find().SetSort(bson.D{{Key: "message_id", Value: 1}}).SetLimit(50)

	// log.Println("Current Message ID:", currentID)
	// log.Println("Filter condition:", filter)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println("Error fetching messages:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message
	for cursor.Next(ctx) {
		var msg domain.Message
		if err := cursor.Decode(&msg); err != nil {
			// log.Println("Error decoding message:", err)
			continue
		}
		messages = append(messages, msg)
	}

	// log.Println("Final messages to be returned:", messages)
	return messages, nil
}
