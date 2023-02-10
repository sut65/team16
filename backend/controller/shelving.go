package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /shelvings

func CreateShelving(c *gin.Context) {

	var shelving entity.Shelving
	var employee entity.Employee
	var label entity.Label
	var stock entity.Stock

	if err := c.ShouldBindJSON(&shelving); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", shelving.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", shelving.Label_ID).First(&label); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "label not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", shelving.Stock_ID).First(&stock); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})
		return
	}

	sv := entity.Shelving{
		Employee: employee,
		Label:    label,
		Stock:    stock,
		Number:   shelving.Number,
		Cost:     shelving.Cost,
	}
	// แทรกการ validate
	if _, err := govalidator.ValidateStruct(&shelving); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&sv).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": shelving})

}

// GET /shelving/:id

func GetShelving(c *gin.Context) {

	var shelving entity.Shelving

	id := c.Param("id")

	if err := entity.DB().Preload("Employee").Preload("Stock").Preload("Label").Raw("SELECT * FROM shelvings WHERE id = ?", id).Find(&shelving).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": shelving})

}

// GET /shelvings

func ListShelvings(c *gin.Context) {

	var shelvings []entity.Shelving

	if err := entity.DB().Preload("Employee").Preload("Stock").Preload("Label").Raw("SELECT * FROM shelvings").Find(&shelvings).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": shelvings})

}

func ListLabel(c *gin.Context) {
	var shelvings []entity.Shelving
	id := c.Param("id")
	if err := entity.DB().Preload("Stock").Preload("Label").Preload("Employee").Raw("SELECT * FROM shelvings WHERE Label_id = ?", id).Find(&shelvings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelvings})
}

// DELETE /shelvings/:id

func DeleteShelving(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM shelvings WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /shelvings

func UpdateShelving(c *gin.Context) {

	var shelving entity.Shelving
	id := c.Param("id")
	var employee entity.Employee
	var label entity.Label
	var stock entity.Stock

	if err := c.ShouldBindJSON(&shelving); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", shelving.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", shelving.Label_ID).First(&label); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "label not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", shelving.Stock_ID).First(&stock); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})
		return
	}

	sv := entity.Shelving{
		Employee: employee,
		Label:    label,
		Stock:    stock,
		Number:   shelving.Number,
		Cost:     shelving.Cost,
	}
	if err := entity.DB().Where("id = ?", id).Updates(&sv).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	// แทรกการ validate
	if _, err := govalidator.ValidateStruct(&sv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shelving})

}
func UpdateNumber(c *gin.Context) {
	var stock entity.Stock
	id := c.Param("id")
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	st := entity.Stock{
		Amount: stock.Amount,
	}

	if err := entity.DB().Where("id = ?", id).Updates(&st).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stock})
}
