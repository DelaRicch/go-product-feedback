package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Comment struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FeedbackId primitive.ObjectID `json:"feedbackId" bson:"feedbackId"`
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Comment string `json:"comment" bson:"comment"`
}

type Feedback struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Details string `json:"details" bson:"details"`
	Upvotes int `json:"upvotes" bson:"upvotes"`
	Status string `json:"status" bson:"status"`
	Comments []Comment `json:"comments" bson:"comments"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}