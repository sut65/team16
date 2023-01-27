package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)


// GET /Employee_attendance
func Listovertime(c *gin.Context) {
	var overtime []entity.Overtime
	if err := entity.DB().Raw("SELECT * FROM overtimes").Scan(&overtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": overtime})
}