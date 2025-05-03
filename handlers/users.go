package handlers

import (
	"bookreviewAPI/database"
	"bookreviewAPI/models"
	"context"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func _isvalidemail(email string) bool{
	_, err:= mail.ParseAddress(email)
	return err == nil
}

func CreateUser(ctx *gin.Context) {
	// Handler to create a new user
	var user models.User
	if err:= 	ctx.ShouldBindJSON(&user); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid Input",
		})
		return
	}
    
	if !_isvalidemail(user.Email){
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"enter valid email address",
		})
		return
	}

	user.ID=primitive.NewObjectID()
	collection:=database.Client.Database("bookreviewsAPI").Collection("users")
	c, cancel:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(c, user)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed to save user to the database",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message":"user registered succesfully",
	})
}