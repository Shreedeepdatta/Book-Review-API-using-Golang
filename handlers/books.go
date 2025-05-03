package handlers

import (
	"bookreviewAPI/database"
	"bookreviewAPI/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(ctx *gin.Context) {
	// Handler to create a new user
	var book models.Book
	if err:= 	ctx.ShouldBindJSON(&book); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid Input",
		})
		return
	}
	book.ID=primitive.NewObjectID()
	collection:=database.Client.Database("bookreviewsAPI").Collection("book")
	c, cancel:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(c, book)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed to save book to the database",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message":"book registered succesfully",
	})
}