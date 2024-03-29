package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /stocks

func CreateStock(c *gin.Context) {

	var stock entity.Stock
	var employee entity.Employee
	var kind entity.Kind
	var storage entity.Storage

	if err := c.ShouldBindJSON(&stock); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", stock.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}

	if tx := entity.DB().Where("id = ?", stock.Kind_ID).First(&kind); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "เลือกชนิดของสินค้าที่ต้องการเพิ่ม"})
		return
	}

	if tx := entity.DB().Where("id = ?", stock.Storage_ID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "เลือกที่จัดเก็บสินค้าที่ที่ต้องการเพิ่ม"})
		return
	}

	// แทรกการ validate
	if _, err := govalidator.ValidateStruct(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	st := entity.Stock{
		Name:     stock.Name,
		Amount:   stock.Amount,
		Price:    stock.Price,
		Employee: employee,
		Kind:     kind,
		Storage:  storage,
		DateTime: stock.DateTime,
	}
	if err := entity.DB().Create(&st).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": st})

}

// GET /stock/:id

func GetStock(c *gin.Context) {
	var stock entity.Stock
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("Kind").Preload("Storage").Raw("SELECT * FROM stocks WHERE id = ?", id).Find(&stock).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stock})
}

// GET /stocks

func ListStocks(c *gin.Context) {

	var stocks []entity.Stock

	if err := entity.DB().Preload("Employee").Preload("Kind").Preload("Storage").Raw("SELECT * FROM stocks").Find(&stocks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stocks})
}

// DELETE /stocks/:id

func DeleteStock(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM stocks WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบสต๊อกที่ต้องการจะลบ"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /stocks

func UpdateStock(c *gin.Context) {

	var stock entity.Stock
	id := c.Param("id")
	var employee entity.Employee
	var kind entity.Kind
	var storage entity.Storage

	if err := c.ShouldBindJSON(&stock); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", stock.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}

	if tx := entity.DB().Where("id = ?", stock.Kind_ID).First(&kind); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "เลือกชนิดสินค้าที่ต้องการอัปเดต"})
		return
	}

	if tx := entity.DB().Where("id = ?", stock.Storage_ID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "เลือกที่จัดเก็บสินค้าที่ต้องการอัปเดต"})
		return
	}

	// แทรกการ validate
	if _, err := govalidator.ValidateStruct(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	st := entity.Stock{
		Name:     stock.Name,
		Amount:   stock.Amount,
		Price:    stock.Price,
		Employee: employee,
		Kind:     kind,
		Storage:  storage,
		DateTime: stock.DateTime,
	}

	if err := entity.DB().Where("id = ?", id).Updates(&st).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": stock})

}
