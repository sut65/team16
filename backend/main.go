package main

import (
<<<<<<< HEAD
	"github.com/Team16/farm_mart/controller"
	"github.com/Team16/farm_mart/entity"
=======
"github.com/Team16/farm_mart/controller"
"github.com/Team16/farm_mart/entity"
>>>>>>> main

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

	// Inventory Routes
	r.GET("/inventories", controller.ListInventories)

	r.GET("/inventory/:id", controller.GetInventory)

	r.POST("/inventories", controller.CreateInventory)

	r.PATCH("/inventories", controller.UpdateInventory)

	r.DELETE("/inventories/:id", controller.DeleteInventory)

	//Stock Routes
	r.GET("/stocks", controller.ListStocks)

	r.GET("/stock/:id", controller.GetStock)

	r.POST("/stocks", controller.CreateStock)

	r.PATCH("/stocks/:id", controller.UpdateStock)

	r.DELETE("/stocks/:id", controller.DeleteStock)

	//Products Routes
	r.GET("/products", controller.ListProducts)

	r.GET("/product/:id", controller.GetProduct)

	r.POST("/products", controller.CreateProduct)

	r.PATCH("/products", controller.UpdateProduct)

	r.DELETE("/products/:id", controller.DeleteProduct)

	//Labels Routes
	r.GET("/labels", controller.ListLabels)

	r.GET("/label/:id", controller.GetLabel)

	r.POST("/labels", controller.CreateLabel)

	r.PATCH("/labels", controller.UpdateLabel)

	r.DELETE("/labels/:id", controller.DeleteLabel)

	//Shelving Routes
	r.GET("/Shelving", controller.ListShelvings)

	r.GET("/Shelving/:id", controller.GetShelving)

	r.POST("/Shelving", controller.CreateShelving)

	r.PATCH("/Shelving", controller.UpdateShelving)

	r.DELETE("/Shelving/:id", controller.DeleteShelving)

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
