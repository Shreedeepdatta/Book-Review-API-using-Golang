package router

import (
	"bookreviewAPI/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine){
	r.GET("/books",handlers.FetchallBooks)
	r.GET("/book/:id", handlers.FetchOneBook)
}