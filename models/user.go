package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reviews struct{
	BookID primitive.ObjectID `json:"book_id" bson:"book_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Review string             `json:"review" bson:"review"`
	Rating int                `json:"rating" bson:"rating"`
}

type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email string                `json:"email" bson:"email"`
	Reviews []Reviews          `json:"reviews" bson:"reviews"`
}

type Book struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
	Author string            `json:"author" bson:"author"`
	ReviewList []Reviews        `json:"reviews" bson:"reviews"`
}
