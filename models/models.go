package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reviews struct {
	BookID primitive.ObjectID `json:"book_id" bson:"book_id" binding:"required"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	Review string             `json:"review" bson:"review" binding:"required,min=1"`
	Rating int                `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
}

type Author struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Name        string               `json:"name" bson:"name" binding:"required"`
	Biography   string               `json:"biography" bson:"biography"`
	Books       []primitive.ObjectID `json:"books" bson:"books"`
	Email       string               `json:"email" bson:"email" binding:"required,email"`
	Nationality string               `json:"nationality" bson:"nationality"`
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username" binding:"required,min=3"`
	Password string             `json:"password" bson:"password" binding:"required,min=6"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	Reviews  []Reviews          `json:"reviews" bson:"reviews" omitempty:"true"`
}

type Book struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title" binding:"required"`
	AuthorID   primitive.ObjectID `json:"author_id" bson:"author_id" binding:"required"`
	Author     string             `json:"author" bson:"author" binding:"required"`
	ReviewList []Reviews          `json:"reviews" bson:"reviews" omitempty:"true"`
}
