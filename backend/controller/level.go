package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /level
func CreateLevel(c *gin.Context) {
	var level entity.Level
	if err := c.ShouldBindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": level})
}

// GET /level/:id
func GetLevel(c *gin.Context) {
	var level entity.Level
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM levels WHERE id = ?", id).Scan(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})
}

// GET /level
func ListLevel(c *gin.Context) {
	var level []entity.Level
	if err := entity.DB().Raw("SELECT * FROM levels").Scan(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})
}

// DELETE /level/:id
func DeleteLevel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM levels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /level
func UpdateLevel(c *gin.Context) {
	var level entity.Level
	if err := c.ShouldBindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", level.ID).First(&level); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
		return
	}

	if err := entity.DB().Save(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})
}