package controller

import (
<<<<<<< HEAD
	"github.com/MaeMethas/se-65-example/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// POST //storages

func CreateShelving(c *gin.Context) {

	var storage entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {
=======
	"net/http"

	"github.com/MaeMethas/se-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /shelvings

func CreateShelving(c *gin.Context) {

	var shelving entity.Shelving

	if err := c.ShouldBindJSON(&shelving); err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	if err := entity.DB().Create(&storage).Error; err != nil {
=======
	if err := entity.DB().Create(&shelving).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": storage})

}

// GET /storage/:id

func GetShelving(c *gin.Context) {

	var storage entity.Storage

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM storages WHERE id = ?", id).Scan(&storage).Error; err != nil {
=======
	c.JSON(http.StatusOK, gin.H{"data": shelving})

}

// GET /shelving/:id

func GetShelving(c *gin.Context) {

	var shelving entity.Shelving

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM shelvings WHERE id = ?", id).Scan(&shelving).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": storage})

}

// GET /storages

func ListShelving(c *gin.Context) {

	var storages []entity.Storage

	if err := entity.DB().Raw("SELECT * FROM storages").Scan(&storages).Error; err != nil {
=======
	c.JSON(http.StatusOK, gin.H{"data": shelving})

}

// GET /users

func ListShelvings(c *gin.Context) {

	var shelvings []entity.Shelving

	if err := entity.DB().Raw("SELECT * FROM shelvings").Scan(&shelvings).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": storages})

}

// DELETE /storages/:id
=======
	c.JSON(http.StatusOK, gin.H{"data": shelvings})

}

// DELETE /shelvings/:id
>>>>>>> main

func DeleteShelving(c *gin.Context) {

	id := c.Param("id")

<<<<<<< HEAD
	if tx := entity.DB().Exec("DELETE FROM storages WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
=======
	if tx := entity.DB().Exec("DELETE FROM shelvings WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})
>>>>>>> main

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

<<<<<<< HEAD
// PATCH /storages

func UpdateShelving(c *gin.Context) {

	var storage entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {
=======
// PATCH /shelvings

func UpdateShelving(c *gin.Context) {

	var shelving entity.Shelving

	if err := c.ShouldBindJSON(&shelving); err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	if tx := entity.DB().Where("id = ?", storage.ID).First(&storage); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
=======
	if tx := entity.DB().Where("id = ?", shelving.ID).First(&shelving); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})
>>>>>>> main

		return

	}

<<<<<<< HEAD
	if err := entity.DB().Save(&storage).Error; err != nil {
=======
	if err := entity.DB().Save(&shelving).Error; err != nil {
>>>>>>> main

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

<<<<<<< HEAD
	c.JSON(http.StatusOK, gin.H{"data": storage})
=======
	c.JSON(http.StatusOK, gin.H{"data": shelving})
>>>>>>> main

}
