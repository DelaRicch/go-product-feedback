package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Details string `json:"details" bson:"details"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}