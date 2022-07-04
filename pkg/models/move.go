package tictactoeapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Move struct {
	Id     primitive.ObjectID `bson:"_id"`
	Row    int                `bson:"row"`
	Column int                `bson:"column"`
}
