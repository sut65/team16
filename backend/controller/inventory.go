package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /inventorys

func CreateInventory(c *gin.Context) {

	var inventory entity.Inventory

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

func GetInventory(c *gin.Context) {

	var inventory entity.Inventory

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM inventories WHERE id = ?", id).Scan(&inventory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})

}

// GET /inventorys

func ListInventories(c *gin.Context) {

	var inventories []entity.Inventory

	if err := entity.DB().Raw("SELECT * FROM inventories").Scan(&inventories).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": inventories})

}

// DELETE /inventories/:id

func DeleteInventory(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM inventories WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /inventories

func UpdateInventory(c *gin.Context) {

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
