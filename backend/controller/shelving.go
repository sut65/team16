package controller

import (
	"github.com/MaeMethas/se-65-example/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST //storages

func CreateShelving(c *gin.Context) {

	var storage entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&storage).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": storage})

}

// GET /storage/:id

func GetShelving(c *gin.Context) {

	var storage entity.Storage

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM storages WHERE id = ?", id).Scan(&storage).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": storage})

}

// GET /storages

func ListShelving(c *gin.Context) {

	var storages []entity.Storage

	if err := entity.DB().Raw("SELECT * FROM storages").Scan(&storages).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": storages})

}

// DELETE /storages/:id

func DeleteShelving(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM storages WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /storages

func UpdateShelving(c *gin.Context) {

	var storage entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", storage.ID).First(&storage); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})

		return

	}

	if err := entity.DB().Save(&storage).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": storage})

}
