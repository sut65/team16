package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /payment_method
func CreatePayment_method(c *gin.Context) {
	var payment_method entity.Payment_method
	if err := c.ShouldBindJSON(&payment_method); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&payment_method).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment_method})
}

// GET /Payment_method/:id
func GetPayment_method(c *gin.Context) {
	var payment_method entity.Payment_method
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM payment_methods WHERE id = ?", id).Scan(&payment_method).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment_method})
}

// GET /Payment_method
func Listpayment_method(c *gin.Context) {
	var payment_method []entity.Payment_method
	if err := entity.DB().Raw("SELECT * FROM payment_methods").Scan(&payment_method).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment_method})
}

// DELETE /Payment_method/:id
func DeletePayment_method(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payment_methods WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_method not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Payment_method
func UpdatePayment_method(c *gin.Context) {
	var payment_method entity.Payment_method
	if err := c.ShouldBindJSON(&payment_method); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", payment_method.ID).First(&payment_method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_method not found"})
		return
	}

	if err := entity.DB().Save(&payment_method).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment_method})
}