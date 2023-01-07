package controller

import (
	"net/http"

<<<<<<< HEAD
	"github.com/MaeMethas/se-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /inventories

func CreateProduct(c *gin.Context) {

	var inventory entity.Product

	if err := c.ShouldBindJSON(&inventory); err != nil {
=======
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /products

func CreateProduct(c *gin.Context) {

	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	if err := entity.DB().Create(&inventory).Error; err != nil {
=======
	if err := entity.DB().Create(&product).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": inventory})

}

// GET /inventory/:id

func GetProduct(c *gin.Context) {

	var inventory entity.Inventory

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM inventorys WHERE id = ?", id).Scan(&inventory).Error; err != nil {
=======
	c.JSON(http.StatusOK, gin.H{"data": product})

}

// GET /product/:id

func GetProduct(c *gin.Context) {

	var product entity.Product

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM products WHERE id = ?", id).Scan(&product).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": inventory})

}

// GET /inventorys

func ListProduct(c *gin.Context) {

	var inventorys []entity.Inventory

	if err := entity.DB().Raw("SELECT * FROM inventorys").Scan(&inventorys).Error; err != nil {
=======
	c.JSON(http.StatusOK, gin.H{"data": product})

}

// GET /products

func ListProducts(c *gin.Context) {

	var products []entity.Product

	if err := entity.DB().Raw("SELECT * FROM products").Scan(&products).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": inventorys})

}

// DELETE /inventorys/:id
=======
	c.JSON(http.StatusOK, gin.H{"data": products})

}

// DELETE /products/:id
>>>>>>> main

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

<<<<<<< HEAD
	if tx := entity.DB().Exec("DELETE FROM inventorys WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})
=======
	if tx := entity.DB().Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
>>>>>>> main

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

<<<<<<< HEAD
// PATCH /inventorys

func UpdateProduct(c *gin.Context) {

	var inventory entity.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
=======
// PATCH /products

func UpdateProduct(c *gin.Context) {

	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	if tx := entity.DB().Where("id = ?", inventory.ID).First(&inventory); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})
=======
	if tx := entity.DB().Where("id = ?", product.ID).First(&product); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
>>>>>>> main

		return

	}

<<<<<<< HEAD
	if err := entity.DB().Save(&inventory).Error; err != nil {
=======
	if err := entity.DB().Save(&product).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": inventory})
=======
	c.JSON(http.StatusOK, gin.H{"data": product})
>>>>>>> main

}
