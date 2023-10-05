package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" validate:"required,min=2,max=100"`
	Complete bool               `json:"iscomplete"`
}
