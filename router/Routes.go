package router

import (
	"bookreviewAPI/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/book", handlers.CreateBook)
	r.POST("/user", handlers.CreateUser)
}
