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
	shelvingID := c.Param("shelvingID")
	var oldShelve entity.Shelving // for check price if discount price more than current price
	var discount entity.Discount
	var shelving entity.Shelving
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
	if tx := entity.DB().Where("id = ?", discount.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Discount_Type_ID).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกปรเภทส่วนลด"})
		return
	}
	if (discount.Discount_s.After(discount.Discount_e)){
		c.JSON(http.StatusBadRequest, gin.H{"error": "วันที่เริ่มลดราคาต้องไม่อยู่หลังจากวันที่วันที่สิ้นสุดการลดราคา"})
		return
	}

	if err := entity.DB().Where("id = ?", shelvingID).First(&oldShelve).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (oldShelve.Cost <= discount.Discount_Price){
		c.JSON(http.StatusBadRequest, gin.H{"error": "ราคาที่ลดต้องไม่ มากกว่าหรือเท่ากับ ราคาของสินค้า"})
		return
	}

	dc := entity.Discount{
		Discount_Price: discount.Discount_Price,             
		Discount_s: discount.Discount_s,
		Discount_e:	discount.Discount_e,             
		Employee:	employee,               
		Discount_Type:	discount_type,  
		Shelving:		shelving,     
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
	if err := entity.DB().Preload("Shelving").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts WHERE id = ?", id).Find(&discount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": discount})
}

// GET /discount
func ListDiscount(c *gin.Context) {
	var discount []entity.Discount
	if err := entity.DB().Preload("Shelving.Stock").Preload("Employee").Preload("Discount_Type").Raw("SELECT * FROM discounts").Find(&discount).Error; err != nil {
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
	id := c.Param("id")
	shelvingID := c.Param("shelvingID")
	var oldShelve entity.Shelving // for check price if discount price more than current price
	var discount entity.Discount
	var shelving entity.Shelving
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
	if tx := entity.DB().Where("id = ?", discount.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกสินค้า"})
		return
	}
	if tx := entity.DB().Where("id = ?", discount.Discount_Type_ID).First(&discount_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกปรเภทส่วนลด"})
		return
	}
	if (discount.Discount_s.After(discount.Discount_e)){
		c.JSON(http.StatusBadRequest, gin.H{"error": "วันที่เริ่มลดราคาต้องไม่อยู่หลังจากวันที่วันที่สิ้นสุดการลดราคา"})
		return
	}

	if err := entity.DB().Where("id = ?", shelvingID).First(&oldShelve).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (oldShelve.Cost <= discount.Discount_Price){
		c.JSON(http.StatusBadRequest, gin.H{"error": "ราคาที่ลดต้องไม่ มากกว่าหรือเท่ากับ ราคาของสินค้า"})
		return
	}

	dc := entity.Discount{
		Discount_Price: discount.Discount_Price,             
		Discount_s: discount.Discount_s,
		Discount_e:	discount.Discount_e,             
		Employee:	employee,               
		Discount_Type:	discount_type,  
		Shelving:		shelving,     
	}
	if _, err := govalidator.ValidateStruct(dc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Where("id = ?", id).Updates(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": discount})
}

// PATCH discounting shelving
func DiscountingShelving(c *gin.Context) {
	var shelving entity.Shelving
	id := c.Param("id")
	if err := c.ShouldBindJSON(&shelving); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sh := entity.Shelving{
		Cost: 	shelving.Cost,
	}
	if err := entity.DB().Where("id = ?", id).Updates(&sh).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": shelving})
}

// PATCH discounting shelving
func ResetCost(c *gin.Context) {
	var shelving entity.Shelving
	shelvingID := c.Param("shelvingID")
	cost, _ := strconv.ParseFloat(c.Param("cost"), 64)

	if err := entity.DB().Where("id = ?", shelvingID).First(&shelving).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newShelving := entity.Shelving{
		Cost: shelving.Cost + cost,
	}

	if err := entity.DB().Model(&shelving).Update("cost", newShelving.Cost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelving})
}