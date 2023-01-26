package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /leave
func CreateLeave(c *gin.Context) {
	var leave entity.Leave
	var employee entity.Employee
	var section entity.Section
	var l_type entity.L_Type

	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.Section_ID).First(&section); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.L_Type_ID).First(&l_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
		return
	}
	lf := entity.Leave{
		Doc_Reason: leave.Doc_Reason,
		Doc_DateS:  leave.Doc_DateS,
		Doc_DateE:  leave.Doc_DateE,
		Doc_Cont:   leave.Doc_Cont,
		Employee:   employee,
		Section:    section,
		L_Type:     l_type,
	}
	if _, err := govalidator.ValidateStruct(lf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&lf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": lf})
}

// GET /leave/:id
func GetLeave(c *gin.Context) {
	var leave entity.Leave
	id := c.Param("id")
	if err := entity.DB().Preload("Section").Preload("L_Type").Preload("Employee").Raw("SELECT * FROM leaves WHERE id = ?", id).Find(&leave).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": leave})
}

// GET /leave
func ListLeave(c *gin.Context) {
	var leave []entity.Leave
	if err := entity.DB().Preload("Section").Preload("L_Type").Preload("Employee").Raw("SELECT * FROM leaves").Find(&leave).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": leave})
}

// DELETE /leave/:id
func DeleteLeave(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM leaves WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "leave not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /leave
func UpdateLeave(c *gin.Context) {
	var leave entity.Leave
	id := c.Param("id")
	var l_type entity.L_Type
	var employee entity.Employee
	var section entity.Section

	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.L_Type_ID).First(&l_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "l_Type not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", leave.Section_ID).First(&section); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "section not found"})
		return
	}
	lf := entity.Leave{
		Doc_Reason: leave.Doc_Reason,             
		Doc_DateS: leave.Doc_DateS,
		Doc_DateE: leave.Doc_DateE,  
		Doc_Cont: leave.Doc_Cont,             
		Employee:	employee,                 
		L_Type:	l_type,  
		Section: section,     
	}
	if err := entity.DB().Where("id = ?", id).Updates(&lf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": leave})
}
