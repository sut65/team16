package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateType_comment(c *gin.Context) {
	var type_commentS entity.Type_Comment
	if err := c.ShouldBindJSON(&type_commentS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&type_commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": type_commentS})
}

// GET /video/:id
func GetType_comment(c *gin.Context) {
	var type_commentS entity.Type_Comment

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM type_comments WHERE id = ?", id).Find(&type_commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_commentS})
}

// GET /videos
func ListType_comments(c *gin.Context) {
	var type_commentS []entity.Reason
	if err := entity.DB().Raw("SELECT * FROM review_points").Find(&type_commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_commentS})
}

// DELETE /videos/:id
func DeleteType_comment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM type_comments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos
func UpdateType_comment(c *gin.Context) {
	var type_commentS entity.Review_Point
	if err := c.ShouldBindJSON(&type_commentS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", type_commentS.ID).First(&type_commentS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review point not found"})
		return
	}

	if err := entity.DB().Save(&type_commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_commentS})
}
