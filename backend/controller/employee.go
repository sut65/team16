package controller

import (
<<<<<<< HEAD
	"github.com/MaeMethas/se-65-example/entity"
=======
	"github.com/Team16/farm_mart/entity"
>>>>>>> main
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /employees

func CreateEmployee(c *gin.Context) {

	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&employee).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employee})

}

// GET /employee/:id

func GetEmployee(c *gin.Context) {

	var employee entity.Employee

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).Scan(&employee).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employee})

}

// GET /employees

func ListEmployees(c *gin.Context) {

	var employees []entity.Employee

	if err := entity.DB().Raw("SELECT * FROM employees").Scan(&employees).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employees})

}

// DELETE /employees/:id

func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /employees

func UpdateEmployee(c *gin.Context) {

	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})

		return

	}

	if err := entity.DB().Save(&employee).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employee})

}
