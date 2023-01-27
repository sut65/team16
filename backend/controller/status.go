package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Status
func CreateStatus(c *gin.Context) {
	var status entity.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET /Status/:id
func GetStatus(c *gin.Context) {
	var status entity.Status
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM statuss WHERE id = ?", id).Scan(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET /Status
func ListStatus(c *gin.Context) {
	var status []entity.Status
	if err := entity.DB().Raw("SELECT * FROM statuses").Scan(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

// DELETE /Status/:id
func DeleteStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Status
func UpdateStatus(c *gin.Context) {
	var status entity.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", status.ID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	if err := entity.DB().Save(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}
