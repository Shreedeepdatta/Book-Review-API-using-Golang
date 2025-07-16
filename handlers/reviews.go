package handlers

import (
	"bookreviewAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReview(ctx *gin.Context){
	var review models.Review
	if err:= ctx.BindJSON(&review); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed to read Body",
		})
		return
	}
	
}