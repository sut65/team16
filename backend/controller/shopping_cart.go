package controller

import (
	"net/http"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Cart
func CreateShopping_Cart(c *gin.Context) {

	var cart entity.Shopping_Cart
	var employee entity.Employee
	var member entity.Member

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Shopping_Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// 10: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", cart.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: ค้นหา Member ด้วย Mem_Tel
	if tx := entity.DB().Where("Mem_Tel = ?", cart.Member_Tal).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}
	// 12: สร้าง Shopping_Cart
	sc := entity.Shopping_Cart{
		Employee:  employee,             // โยงความสัมพันธ์กับ Entity Employee
		Member:    member,               // โยงความสัมพันธ์กับ Entity Member
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cart})
}


// GET /Shopping_Cart
func ListShopping_Cart(c *gin.Context) {
	var cart []entity.Shopping_Cart
	if err := entity.DB().Raw("SELECT * FROM carts").Scan(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DELETE /Shopping_Cart/:id
func DeleteShopping_Cart(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM carts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Shopping_Cart
func UpdateShopping_Cart(c *gin.Context) {
	var cart entity.Shopping_Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cart.ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	if err := entity.DB().Save(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}