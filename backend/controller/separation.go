package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateSeparation(c *gin.Context) {

	var separation entity.Separation
	var reason entity.Reason
	var employee entity.Employee
	var shelving entity.Shelving

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&separation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Reason_ID).First(&reason); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reason not found"})
		return
	}
	// 12: สร้าง WatchVideo
	wv := entity.Separation{
		Amount:      separation.Amount,
		Status:      separation.Status,
		Reason:      reason,                 // โยงความสัมพันธ์กับ Entity Resolution
		Shelving:    shelving,               // โยงความสัมพันธ์กับ Entity Video
		Employee:    employee,               // โยงความสัมพันธ์กับ Entity Playlist
		Date_Out:    separation.Date_Out,    // ตั้งค่าฟิลด์ watchedTime
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /watchvideo/:id
func GetSeparation(c *gin.Context) {
	var separationS entity.Separation
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("Reason").Preload("Shelving").Raw("SELECT * FROM separations WHERE id = ?", id).Find(&separationS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": separationS})
}

// GET /watch_videos
func ListSeparations(c *gin.Context) {
	var separationS []entity.Separation
	if err := entity.DB().Preload("Employee").Preload("Reason").Preload("Shelving").Raw("SELECT * FROM separations").Find(&separationS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": separationS})
}

// DELETE /watch_videos/:id
func DeleteSeparation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM separations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "separation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateSeparation(c *gin.Context) {

	var separationS entity.Separation
	id := c.Param("id")
	var reason entity.Reason
	var employee entity.Employee
	var shelving entity.Shelving

	if err := c.ShouldBindJSON(&separationS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Reason_ID).First(&reason); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reason not found"})
		return
	}
	dc := entity.Separation{           
		Date_Out:       separationS.Date_Out,
		Employee:	    separationS.Employee,               
		Shelving:	    shelving,               
		Reason:	    	reason,               
		Amount:	        separationS.Amount,  
		Status:		    separationS.Status,     
	}
	if err := entity.DB().Where("id = ?", id).Updates(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": separationS})


}
