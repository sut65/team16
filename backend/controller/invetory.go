package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.con/MaeMethas/se-65-example/entity"
)

// POST /invetorys

func CreateInvetory(c *gin.Context) {

	var invetory entity.Invetory

	if err := c.ShouldBindJSON(&invetory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&invetory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": invetory})

}

// GET /inventory/:id

func GetInvetory(c *gin.Context) {

	var invetory entity.Invetory

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM invetorys WHERE id = ?", id).Scan(&invetory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": invetory})

}

// GET /inventorys

func ListInvetorys(c *gin.Context) {

	var invetorys []entity.Invetory

	if err := entity.DB().Raw("SELECT * FROM invetorys").Scan(&invetorys).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": invetorys})

}

// DELETE /invetorys/:id

func DeleteInvetory(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM invetorys WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /invetorys

func UpdateInvetory(c *gin.Context) {

	var invetory entity.Invetory

	if err := c.ShouldBindJSON(&invetory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", invetory.ID).First(&invetory); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invetory not found"})

		return

	}

	if err := entity.DB().Save(&invetory).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": invetory})

}
