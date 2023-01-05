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

	// Employee Routes

	r.GET("/emloyees", controller.ListEmployees)

	r.GET("/emloyee/:id", controller.GetEmployee)

	r.POST("/emloyees", controller.CreateEmployee)

	r.PATCH("/emloyees", controller.UpdateEmployee)

	r.DELETE("/emloyees/:id", controller.DeleteEmployee)

	//Kind Routes
	r.GET("/kinds", controller.ListKinds)

	r.GET("/kind/:id", controller.GetKind)

	r.POST("/kinds", controller.CreateKind)

	r.PATCH("/kinds", controller.UpdateKind)

	r.DELETE("/kinds/:id", controller.DeleteKind)

	// Product Routes
	r.GET("/products", controller.ListProducts)

	r.GET("/product/:id", controller.GetProduct)

	r.POST("/products", controller.CreateProduct)

	r.PATCH("/products/:id", controller.UpdateProduct)

	r.DELETE("/products/:id", controller.DeleteProduct)

	//Stock Routes
	r.GET("/stocks", controller.ListStocks)

	r.GET("/stock/:id", controller.GetStock)

	r.POST("/stocks", controller.CreateStock)

	r.PATCH("/stocks/:id", controller.UpdateStock)

	r.DELETE("/stocks/:id", controller.DeleteStock)

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
