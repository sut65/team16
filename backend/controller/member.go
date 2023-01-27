package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /member
func CreateMember(c *gin.Context) {
	var member entity.Member
	var employee entity.Employee
	var gender entity.Gender
	var level entity.Level

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Level_ID).First(&level); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
		return
	}
	mb := entity.Member{
		Mem_Name: member.Mem_Name,
		Mem_Age:  member.Mem_Age,
		Mem_Tel:  member.Mem_Tel,
		Employee: employee,
		Gender:   gender,
		Level:    level,
	}
	if _, err := govalidator.ValidateStruct(mb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&mb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": mb})
}

// GET /member/:id
func GetMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	if err := entity.DB().Preload("Gender").Preload("Level").Preload("Employee").Raw("SELECT * FROM members WHERE id = ?", id).Find(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /member
func ListMember(c *gin.Context) {
	var member []entity.Member
	if err := entity.DB().Preload("Gender").Preload("Level").Preload("Employee").Raw("SELECT * FROM members").Find(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// DELETE /member/:id
func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM members WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /member
func UpdateMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	var gender entity.Gender
	var employee entity.Employee
	var level entity.Level

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", member.Level_ID).First(&level); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
		return
	}
	lf := entity.Member{
		Mem_Name: member.Mem_Name,             
		Mem_Age: member.Mem_Age,
		Mem_Tel: member.Mem_Tel,             
		Employee: employee,                
		Gender:	gender,  
		Level: level,     
	}
	if err := entity.DB().Where("id = ?", id).Updates(&lf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}
