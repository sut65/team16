package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /l_type
func CreateL_Type(c *gin.Context) {
	var l_type entity.L_Type
	if err := c.ShouldBindJSON(&l_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&l_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": l_type})
}

// GET /l_type/:id
func GetL_Type(c *gin.Context) {
	var l_type entity.L_Type
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM l_types WHERE id = ?", id).Scan(&l_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": l_type})
}

// GET /l_type
func ListL_Type(c *gin.Context) {
	var l_type []entity.L_Type
	if err := entity.DB().Raw("SELECT * FROM l_types").Scan(&l_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": l_type})
}

// DELETE /l_type/:id
func DeleteL_Type(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM l_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "l_type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /l_type
func UpdateL_Type(c *gin.Context) {
	var l_type entity.L_Type
	if err := c.ShouldBindJSON(&l_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", l_type.ID).First(&l_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "l_type not found"})
		return
	}

	if err := entity.DB().Save(&l_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": l_type})
}