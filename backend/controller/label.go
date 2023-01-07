package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Labels

func CreateLabel(c *gin.Context) {

	var label entity.Label

	if err := c.ShouldBindJSON(&label); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&label).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": label})

}

// GET /Label/:id

func GetLabel(c *gin.Context) {

	var label entity.Label

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM labels WHERE id = ?", id).Scan(&label).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": label})

}

// GET /labels

func ListLabels(c *gin.Context) {

	var labels []entity.Label

	if err := entity.DB().Raw("SELECT * FROM labels").Scan(&labels).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": labels})

}

// DELETE /Labels/:id

func DeleteLabel(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM labels WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Label not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Labels

func UpdateLabel(c *gin.Context) {

	var Label entity.Label

	if err := c.ShouldBindJSON(&Label); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Label.ID).First(&Label); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Label not found"})

		return

	}

	if err := entity.DB().Save(&Label).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Label})

}
