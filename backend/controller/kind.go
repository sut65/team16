package controller

import (
	"github.com/gin-gonic/gin"
	"github.con/MaeMethas/se-65-example/entity"
	"net/http"
)

// POST /kinds

func CreateKind(c *gin.Context) {

	var kind entity.Kind

	if err := c.ShouldBindJSON(&kind); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&kind).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": kind})

}

// GET /kind/:id

func GetKind(c *gin.Context) {

	var kind entity.Kind

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM kinds WHERE id = ?", id).Scan(&kind).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": kind})

}

// GET /users

func ListKinds(c *gin.Context) {

	var kinds []entity.Kind

	if err := entity.DB().Raw("SELECT * FROM kinds").Scan(&kinds).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": kinds})

}

// DELETE /kinds/:id

func DeleteKind(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM kinds WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "kind not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /kinds

func UpdateKind(c *gin.Context) {

	var kind entity.Kind

	if err := c.ShouldBindJSON(&kind); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", kind.ID).First(&kind); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "kind not found"})

		return

	}

	if err := entity.DB().Save(&kind).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": kind})

}
