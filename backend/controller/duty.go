package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)


// GET /Employee_attendance
func Listduty(c *gin.Context) {
	var duty []entity.Duty
	if err := entity.DB().Raw("SELECT * FROM duties").Scan(&duty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": duty})
}