package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /gender
func CreateGender(c *gin.Context) {
	var gender entity.Gender
	if err := c.ShouldBindJSON(&gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gender})
}

// GET /gender/:id
func GetGender(c *gin.Context) {
	var gender entity.Gender
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genders WHERE id = ?", id).Scan(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gender})
}

// GET /gender
func ListGender(c *gin.Context) {
	var gender []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gender})
}

// DELETE /gender/:id
func DeleteGender(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM genders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /gender
func UpdateGender(c *gin.Context) {
	var gender entity.Gender
	if err := c.ShouldBindJSON(&gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", gender.ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	if err := entity.DB().Save(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gender})
}