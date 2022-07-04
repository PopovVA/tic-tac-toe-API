package tictactoeapi

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lobby struct {
	Id        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
}
