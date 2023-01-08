package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateReview_point(c *gin.Context) {
	var review_pointS entity.Review_Point
	if err := c.ShouldBindJSON(&review_pointS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&review_pointS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review_pointS})
}

// GET /video/:id
func GetReview_point(c *gin.Context) {
	var review_pointS entity.Review_Point

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM review_points WHERE id = ?", id).Find(&review_pointS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review_pointS})
}

// GET /videos
func ListReview_points(c *gin.Context) {
	var review_pointS []entity.Reason
	if err := entity.DB().Raw("SELECT * FROM review_points").Find(&review_pointS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review_pointS})
}

// DELETE /videos/:id
func DeleteReview_point(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM review_points WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review point not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos
func UpdateReview_point(c *gin.Context) {
	var review_pointS entity.Review_Point
	if err := c.ShouldBindJSON(&review_pointS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", review_pointS.ID).First(&review_pointS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review point not found"})
		return
	}

	if err := entity.DB().Save(&review_pointS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review_pointS})
}
