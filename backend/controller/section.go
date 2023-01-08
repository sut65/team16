package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /section
func CreateSection(c *gin.Context) {
	var section entity.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&section).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": section})
}

// GET /section/:id
func GetSection(c *gin.Context) {
	var section entity.Section
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM sections WHERE id = ?", id).Scan(&section).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": section})
}

// GET /section
func ListSection(c *gin.Context) {
	var section []entity.Section
	if err := entity.DB().Raw("SELECT * FROM sections").Scan(&section).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": section})
}

// DELETE /section/:id
func DeleteSection(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sections WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "section not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /section
func UpdateSection(c *gin.Context) {
	var section entity.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", section.ID).First(&section); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "section not found"})
		return
	}

	if err := entity.DB().Save(&section).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": section})
}