package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateReason(c *gin.Context) {
	var reasonS entity.Reason
	if err := c.ShouldBindJSON(&reasonS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&reasonS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reasonS})
}

// GET /video/:id
func GetReason(c *gin.Context) {
	var reasonS entity.Reason

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM reasons WHERE id = ?", id).Find(&reasonS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reasonS})
}

// GET /videos
func ListReasons(c *gin.Context) {
	var reasonS []entity.Reason
	if err := entity.DB().Raw("SELECT * FROM reasons").Find(&reasonS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reasonS})
}

// DELETE /videos/:id
func DeleteReason(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM reasons WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reason not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos
func UpdateVideo(c *gin.Context) {
	var reasonS entity.Reason
	if err := c.ShouldBindJSON(&reasonS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", reasonS.ID).First(&reasonS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reason not found"})
		return
	}

	if err := entity.DB().Save(&reasonS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reasonS})
}
