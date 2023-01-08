package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateComment(c *gin.Context) {

	var comment entity.Comment
	var review_point entity.Review_Point
	var type_comment entity.Type_Comment
	var payment entity.Payment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Review_point_ID).First(&review_point); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review point not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Type_Com_ID).First(&type_comment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type comment not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Payment_ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}
	// 12: สร้าง WatchVideo
	wv := entity.Comment{
		Review_point:      review_point,                 // โยงความสัมพันธ์กับ Entity Resolution
		Payment:           payment,               // โยงความสัมพันธ์กับ Entity Video
		Type_Com:          type_comment,               // โยงความสัมพันธ์กับ Entity Playlist
		Date_Now:          comment.Date_Now,    // ตั้งค่าฟิลด์ watchedTime
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
func GetComment(c *gin.Context) {
	var commentS entity.Comment
	id := c.Param("id")
	if err := entity.DB().Preload("Review_point").Preload("Payment").Preload("Type_Com").Raw("SELECT * FROM comments WHERE id = ?", id).Find(&commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": commentS})
}

// GET /watch_videos
func ListComments(c *gin.Context) {
	var commentS []entity.Comment
	if err := entity.DB().Preload("Review_point").Preload("Payment").Preload("Type_Com").Raw("SELECT * FROM comments").Find(&commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentS})
}

// DELETE /watch_videos/:id
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM comments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateComment(c *gin.Context) {
	var commentS entity.Comment
	if err := c.ShouldBindJSON(&commentS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", commentS.ID).First(&commentS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment not found"})
		return
	}

	if err := entity.DB().Save(&commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentS})
}
