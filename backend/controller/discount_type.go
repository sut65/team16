package controller

import (
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /discount_type

func CreateDiscount_Type(c *gin.Context) {

	var discount_type entity.Discount_Type
	if err := c.ShouldBindJSON(&discount_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&discount_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount_type})

}

// GET /discount_type/:id

func GetDiscount_Type(c *gin.Context) {

	var discount_type entity.Discount_Type
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM discount_types WHERE id = ?", id).Scan(&discount_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount_type})

}

// GET /discount_types

func ListDiscount_Type(c *gin.Context) {
	var discount_type []entity.Discount_Type
	if err := entity.DB().Raw("SELECT * FROM discount_types").Scan(&discount_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount_type})
}

// DELETE /discount_type/:id

func DeleteDiscount_Type(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM discount_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /discount_types

func UpdateDiscount_Type(c *gin.Context) {
	var discount_type entity.Discount_Type
	if err := c.ShouldBindJSON(&discount_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", discount_type.ID).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}

	if err := entity.DB().Save(&discount_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount_type})

}
