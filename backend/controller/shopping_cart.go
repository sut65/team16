package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Cart
func CreateShopping_Cart(c *gin.Context) {

	var Shopping_Cart entity.Shopping_Cart
	var resolution entity.Resolution
	var playlist entity.Playlist
	var video entity.Video

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Shopping_Cart
	if err := c.ShouldBindJSON(&Shopping_Cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", Shopping_Cart.VideoID).First(&video); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", Shopping_Cart.ResolutionID).First(&resolution); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 11: ค้นหา playlist ด้วย id
	if tx := entity.DB().Where("id = ?", Shopping_Cart.PlaylistID).First(&playlist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}
	// 12: สร้าง Shopping_Cart
	wv := entity.Shopping_Cart{
		Resolution:  resolution,             // โยงความสัมพันธ์กับ Entity Resolution
		Video:       video,                  // โยงความสัมพันธ์กับ Entity Video
		Playlist:    playlist,               // โยงความสัมพันธ์กับ Entity Playlist
		WatchedTime: Shopping_Cart.WatchedTime, // ตั้งค่าฟิลด์ watchedTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(