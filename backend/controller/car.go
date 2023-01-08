package controller

import (
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /car

func CreateCar(c *gin.Context) {

	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})

}

// GET /car/:id

func GetCar(c *gin.Context) {
	var car entity.Car
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM cars WHERE id = ?", id).Scan(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})

}

// GET /cars

func ListCar(c *gin.Context) {
	var car []entity.Car
	if err := entity.DB().Raw("SELECT * FROM cars").Scan(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

// DELETE /car/:id

func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cars WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /cars

func UpdateCar(c *gin.Context) {
	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", car.ID).First(&car); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}

	if err := entity.DB().Save(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})

}
