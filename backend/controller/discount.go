package controller

import (
	"net/http"
	"strconv"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /discount
func CreateDiscount(c *gin.Context) {
	var discount entity.Discount
	var inventory entity.Stock
	var employee entity.Employee
	var discount_type entity.Discount_Type

	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกพนักงาน"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Stock_ID).First(&inventory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Discount_Type_ID).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกปรเภทส่วนลด"})
		return
	}
	dc := entity.Discount{
		Discount_Price: discount.Discount_Price,             
		Discount_s: discount.Discount_s,
		Discount_e:	discount.Discount_e,             
		Employee:	employee,               
		Discount_Type:	discount_type,  
		Stock:		inventory,     
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
	if err := entity.DB().Preload("Stock").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts WHERE id = ?", id).Find(&discount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount})
}

// GET /discount
func ListDiscount(c *gin.Context) {
	var discount []entity.Discount
	if err := entity.DB().Preload("Stock").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts").Find(&discount).Error; err != nil {
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
	id := c.Param("id")
	var inventory entity.Stock
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
	if tx := entity.DB().Where("id = ?", discount.Stock_ID).First(&inventory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Discount_Type_ID).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}
	dc := entity.Discount{
		Discount_Price: discount.Discount_Price,             
		Discount_s: discount.Discount_s,
		Discount_e:	discount.Discount_e,             
		Employee:	employee,               
		Discount_Type:	discount_type,  
		Stock:		inventory,     
	}
	if err := entity.DB().Where("id = ?", id).Updates(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": discount})
}

// PATCH discounting stock
func DiscountingStock(c *gin.Context) {
	var stock entity.Stock
	id := c.Param("id")
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	st := entity.Stock{
		Price: 	stock.Price,
	}
	if err := entity.DB().Where("id = ?", id).Updates(&st).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stock})
}

// PATCH discounting stock
func ResetPrice(c *gin.Context) {
	var stock entity.Stock
	stockID := c.Param("stockID")
	price, _ := strconv.ParseFloat(c.Param("price"), 64)

	if err := entity.DB().Where("id = ?", stockID).First(&stock).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStock := entity.Stock{
		Price: stock.Price + price,
	}

	if err := entity.DB().Model(&stock).Update("price", newStock.Price).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stock})
}