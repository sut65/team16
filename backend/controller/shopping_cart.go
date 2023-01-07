package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Carts

func CreateShopping_Cart(c *gin.Context) {

	var Shopping_Cart entity.Shopping_Cart

	if err := c.ShouldBindJSON(&Shopping_Cart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Shopping_Cart).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Shopping_Cart})

}

// GET /Shopping_Cart/:id

func GetShopping_Cart(c *gin.Context) {

	var Shopping_Cart entity.Shopping_Cart

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Shopping_Carts WHERE id = ?", id).Scan(&Shopping_Cart).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Shopping_Cart})

}

// GET /users

func ListShopping_Carts(c *gin.Context) {

	var Shopping_Carts []entity.Shopping_Cart

	if err := entity.DB().Raw("SELECT * FROM Shopping_Carts").Scan(&Shopping_Carts).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Shopping_Carts})

}

// DELETE /Shopping_Carts/:id

func DeleteShopping_Cart(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Shopping_Carts WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Shopping_Cart not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Shopping_Carts

func UpdateShopping_Cart(c *gin.Context) {

	var Shopping_Cart entity.Shopping_Cart

	if err := c.ShouldBindJSON(&Shopping_Cart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", Shopping_Cart.ID).First(&Shopping_Cart); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Shopping_Cart not found"})

		return

	}

	if err := entity.DB().Save(&Shopping_Cart).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Shopping_Cart})

}
