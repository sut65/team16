package controller

import (
	"net/http"

	"github.com/MaeMethas/se-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /inventories

func CreateProduct(c *gin.Context) {

	var inventory entity.Product

	if err := c.ShouldBindJSON(&inventory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&inventory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})

}

// GET /inventory/:id

func GetProduct(c *gin.Context) {

	var inventory entity.Inventory

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM inventorys WHERE id = ?", id).Scan(&inventory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})

}

// GET /inventorys

func ListProduct(c *gin.Context) {

	var inventorys []entity.Inventory

	if err := entity.DB().Raw("SELECT * FROM inventorys").Scan(&inventorys).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventorys})

}

// DELETE /inventorys/:id

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM inventorys WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /inventorys

func UpdateProduct(c *gin.Context) {

	var inventory entity.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", inventory.ID).First(&inventory); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})

		return

	}

	if err := entity.DB().Save(&inventory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})

}
