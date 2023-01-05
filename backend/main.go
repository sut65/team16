package main

import (
	"github.con/MaeMethas/se-65-example/controller"
	"github.con/MaeMethas/se-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// User Routes

	r.GET("/users", controller.ListEmployees)

	r.GET("/user/:id", controller.GetEmployee)

	r.POST("/users", controller.CreateEmployee)

	r.PATCH("/users", controller.UpdateEmployee)

	r.DELETE("/users/:id", controller.DeleteEmployee)

	// Run the server

	r.Run()

}
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
