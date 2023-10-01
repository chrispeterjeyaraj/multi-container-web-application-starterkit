package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required,min=2,max=100"`
	IsComplete bool               `json:"iscomplete"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
