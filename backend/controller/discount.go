package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /discount
func CreateDiscount(c *gin.Context) {
	var discount entity.Discount
	var inventory entity.Inventory
	var employee entity.Employee
	var discount_type entity.Discount_Type

	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Discount_Type).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Inventory_ID).First(&inventory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})
		return
	}
	dc := entity.Discount{
		Discount_Price: discount.Discount_Price,             
		Discount_s: discount.Discount_s,
		Discount_e:	discount.Discount_e,             
		Employee:	employee,               
		Discount_Type:	discount_type,  
		Inventory:		inventory,     
	}
	if _, err := govalidator.ValidateStruct(dc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"data": dc})
}

// GET /discount/:id
func GetDiscount(c *gin.Context) {
	var discount entity.Discount
	id := c.Param("id")
	if err := entity.DB().Preload("Inventory").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts WHERE id = ?", id).Find(&discount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount})
}

// GET /discount
func ListDiscount(c *gin.Context) {
	var discount []entity.Discount
	if err := entity.DB().Preload("Inventory").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts").Find(&discount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": discount})
}

// DELETE /discount/:id
func DeleteDiscount(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM discounts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this discount not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /discount
func UpdateDiscount(c *gin.Context) {
	var discount entity.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", discount.ID).First(&discount); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this discount not found"})
		return
	}

	if err := entity.DB().Save(&discount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": discount})
}