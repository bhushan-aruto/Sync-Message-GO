package domain

type Message struct {
	MessageId     string `bson:"message_id" json:"message_id"`
	PrevMessageId string `bson:"prev_message_id" json:"prev_message_id"`
	MessageFrom   string `bson:"message_from" json:"message_from"`

	MessageTo      string `bson:"message_to" json:"message_to"`
	MessageContent string `bson:"message_content" json:"message_content"`
	MessageStatus  string `bson:"message_status" json:"message_status"`
	CreatedAt      string `bson:"created_at" json:"created_at"`
	DeliveredAt    string `bson:"delivered_at" json:"delivered_at"`
	ReadAt         string `bson:"read_at" json:"read_at"`
}

type Messagerepo interface {
	GetOrderdMessage(userId string, currentMessageid string) ([]Message, error)
}
