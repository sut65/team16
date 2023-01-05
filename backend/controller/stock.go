package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.con/MaeMethas/se-65-example/entity"
)

// POST /stocks

func CreateStock(c *gin.Context) {

	var stock entity.Stock

	if err := c.ShouldBindJSON(&stock); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&stock).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": stock})

}

// GET /stock/:id

func GetStock(c *gin.Context) {

	var stock entity.Stock

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM stocks WHERE id = ?", id).Scan(&stock).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": stock})

}

// GET /users

func ListStocks(c *gin.Context) {

	var stocks []entity.Stock

	if err := entity.DB().Raw("SELECT * FROM stocks").Scan(&stocks).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": stocks})

}

// DELETE /stocks/:id

func DeleteStock(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM stocks WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /stocks

func UpdateStock(c *gin.Context) {

	var stock entity.Stock

	if err := c.ShouldBindJSON(&stock); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", stock.ID).First(&stock); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "stock not found"})

		return

	}

	if err := entity.DB().Save(&stock).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": stock})

}
