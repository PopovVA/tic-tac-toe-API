package tictactoeapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Move struct {
	Id       primitive.ObjectID `bson:"_id"`
	PlayedId Player             `bson:"playerId"`
	GameId   Game               `bson:"gameId"`
	Row      int                `bson:"rowNumber"`
	Column   int                `bson:"columnNumber"`
	MovedAt  int                `bson:"movedAt"`
}
