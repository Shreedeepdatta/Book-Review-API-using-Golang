package handlers

import "github.com/gin-gonic/gin"

func FetchallBooks(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "router for fetching all books",
	})
}

func FetchOneBook(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"message": id,
	})
}
