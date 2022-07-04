package tictactoeapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type History struct {
	Id    primitive.ObjectID `bson:"_id"`
	Game  Game               `bson:"game"`
	Moves []Move             `bson:"moves"`
}
