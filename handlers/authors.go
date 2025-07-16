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

func CreateAuthor(ctx *gin.Context){
	var author models.Author
	if err:=ctx.BindJSON(&author); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid input",
		})
	    return
	}
	author.ID=primitive.NewObjectID()
	collection:=database.Client.Database("bookreviewsAPI").Collection("author")
	c, cancel:=context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_,err:=collection.InsertOne(c, author)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"error registering the author",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message":"author registered successfully",
	})
}