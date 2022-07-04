package tictactoeapi

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"userName"`
}
