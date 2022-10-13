package main

import (
	"github.com/B6304416/sa-65-example/controller"
	"github.com/B6304416/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

}

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Activity Routes
	r.GET("/Activitys", controller.ListActivity)
	r.GET("/Activity/:id", controller.GetActivity)
	r.POST("/Activitys", controller.CreateActivity)
	r.PATCH("/Activitys", controller.UpdateActivity)
	r.DELETE("/Activitys/:id", controller.DeleteActivity)
	// Run the server
	r.Run()

}
