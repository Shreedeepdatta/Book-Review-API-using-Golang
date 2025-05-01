package main

import (
	"bookreviewAPI/database"
	"bookreviewAPI/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMongo()
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "welcome to book review API",
		})
	})
	router.RegisterRoutes(r)
	r.Run(":9000")
}
