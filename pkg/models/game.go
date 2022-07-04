package tictactoeapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id     primitive.ObjectID `bson:"_id"`
	Winner Player             `bson:"winner"`
}
