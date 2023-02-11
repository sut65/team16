package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /delivery
func CreateDelivery(c *gin.Context) {
	var delivery entity.Delivery
	var car entity.Car
	var employee entity.Employee
	var payment entity.Payment
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกพนักงานที่ส่งสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Car_ID).First(&car); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกรถที่ใช้ส่งสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Payment_ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกรายการสินค้าที่ส่ง"})
		return
	}
	dl := entity.Delivery{
		Location: delivery.Location,             
		Customer_name:	delivery.Customer_name,
		Delivery_date:	delivery.Delivery_date,             
		Employee:	employee,               
		Car:		car,  
		Payment:	payment,     
	}
	if _, err := govalidator.ValidateStruct(dl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&dl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"data": dl})
}

// GET /delivery/:id
func GetDelivery(c *gin.Context) {
	var delivery entity.Delivery
	id := c.Param("id")
	if err := entity.DB().Preload("Car").Preload("Employee").Preload("Payment").Raw("SELECT * FROM deliveries WHERE id = ?", id).Find(&delivery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": delivery})
}

// GET /delivery
func ListDelivery(c *gin.Context) {
	var delivery []entity.Delivery
	if err := entity.DB().Preload("Car").Preload("Employee").Preload("Payment").Raw("SELECT * FROM deliveries").Find(&delivery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": delivery})
}

// DELETE /delivery/:id
func DeleteDelivery(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM deliveries WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this delivery not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /delivery
func UpdateDelivery(c *gin.Context) {
	var delivery entity.Delivery
	id := c.Param("id")
	var car entity.Car
	var employee entity.Employee
	var payment entity.Payment

	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกพนักงานที่ส่งสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Car_ID).First(&car); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกรถที่ใช้ส่งสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", delivery.Payment_ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกรายการสินค้าที่ส่ง"})
		return
	}
	dl := entity.Delivery{
		Location: delivery.Location,             
		Customer_name:	delivery.Customer_name,
		Delivery_date:	delivery.Delivery_date,             
		Employee:	employee,               
		Car:		car,  
		Payment:	payment,     
	}
	if _, err := govalidator.ValidateStruct(dl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Where("id = ?", id).Updates(&dl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": delivery})
}