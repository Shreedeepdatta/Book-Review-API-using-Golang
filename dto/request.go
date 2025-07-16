package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required,min=3"`
	Password  string             `json:"password" bson:"password" binding:"required,min=6"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	IsAuthor  bool               `json:"is_author" bson:"is_author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type Author struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Biography   string             `json:"biography" bson:"biography"`
	Email       string             `json:"email" bson:"email" binding:"required,email"`
	Nationality string             `json:"nationality" bson:"nationality"`
	BooksCount  int                `json:"books_count" bson:"books_count"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type Book struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	AuthorID    primitive.ObjectID `json:"author_id" bson:"author_id" binding:"required"`
	AuthorName  string             `json:"author_name" bson:"author_name"`
	Description string             `json:"description" bson:"description"`
	Genre       string             `json:"genre" bson:"genre"`
	ISBN        string             `json:"isbn" bson:"isbn"`
	PublishedAt time.Time          `json:"published_at" bson:"published_at"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ReviewCount int                `json:"review_count" bson:"review_count"`
	AvgRating   float64            `json:"avg_rating" bson:"avg_rating"`
}

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BookID    primitive.ObjectID `json:"book_id" bson:"book_id" binding:"required"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	Username  string             `json:"username" bson:"username"`
	Review    string             `json:"review" bson:"review" binding:"required,min=1"`
	Rating    int                `json:"rating" bson:"rating" binding:"required,min=1,max=5"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// Request structs
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	IsAuthor bool   `json:"is_author"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Biography   string `json:"biography"`
	Nationality string `json:"nationality"`
}

type CreateBookRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	ISBN        string    `json:"isbn"`
	PublishedAt time.Time `json:"published_at"`
}

type CreateReviewRequest struct {
	Rating int    `json:"rating" binding:"required,min=1,max=5"`
	Review string `json:"review" binding:"required,min=1"`
}
