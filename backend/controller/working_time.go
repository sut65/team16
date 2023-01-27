package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)


// GET /Employee_attendance
func Listworking_time(c *gin.Context) {
	var working_time []entity.Working_time
	if err := entity.DB().Raw("SELECT * FROM working_times").Scan(&working_time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": working_time})
}