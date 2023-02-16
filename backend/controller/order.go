package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Order
func CreateOrder(c *gin.Context) {

	var order entity.Order
	var shelv entity.Shelving
	var cart entity.Shopping_Cart

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 10: ค้นหา cart ด้วย id
	if tx := entity.DB().Where("id = ?", order.Shopping_Cart_ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบตะกร้า"})
		return
	}

	// 11: ค้นหา shelv ด้วย id
	if tx := entity.DB().Where("id = ?", order.Shelving_ID).First(&shelv); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบสินค้า"})
		return
	}
	// 12: สร้าง Order
	sc := entity.Order{
		Quantity:      order.Quantity,
		Prices:        order.Prices,
		Shopping_Cart: cart,
		Shelving:      shelv, // โยงความสัมพันธ์กับ Entity shelving
	}

	if _, err := govalidator.ValidateStruct(sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// GET /Order/:id
func GetOrder(c *gin.Context) {
	var order entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM orders WHERE id = ?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Shelving_ID": order.Shelving_ID,
		"Shelving_Number": order.Shelving.Number,
	 })
}

// GET /Order
func ListOrder(c *gin.Context) {
	var order []entity.Order
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM orders").Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
func ListShelv(c *gin.Context) {
	var shelvings []entity.Shelving
	if err := entity.DB().Preload("Stock").Raw("SELECT * FROM shelvings").Find(&shelvings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": shelvings})
}

// ListOrderCart /OrderCart/:id
func ListOrderCart(c *gin.Context) {
	var order []entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("Shelving").Preload("Shelving.Stock").Preload("Shopping_Cart").Raw("SELECT * FROM orders WHERE shopping_cart_id = ?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func ListOrdersum(c *gin.Context) {
	var order []entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM orders WHERE shopping_cart_id = ?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sumPrices float64
	for _, o := range order {
		sumPrices += o.Prices
	}

	c.JSON(http.StatusOK, gin.H{"sumPrices": sumPrices})
}

// DELETE /Order/:id
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Order
func UpdateOrder(c *gin.Context) {
	var order entity.Order
	id := c.Param("id")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	od := entity.Order{
		Prices: order.Prices,
		Quantity: order.Quantity,
	}
	if _, err := govalidator.ValidateStruct(od); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Where("id = ?", id).Updates(&od).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
func UpdateQuantity(c *gin.Context) {
	var shelv entity.Shelving
	id := c.Param("id")
	if err := c.ShouldBindJSON(&shelv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sv := entity.Shelving{
		Number: shelv.Number,
		//Cost: shelv.Cost,
	}
	// if _, err := govalidator.ValidateStruct(sv); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := entity.DB().Where("id = ?", id).Updates(&sv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": shelv})
}
