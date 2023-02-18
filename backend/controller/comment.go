package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /comment
func CreateComment(c *gin.Context) {

	var comment entity.Comment
	var review_point entity.Review_Point
	var type_comment entity.Type_Comment
	var payment entity.Payment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 6 จะถูก bind เข้าตัวแปร comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 7: ค้นหา type_comment ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Type_Com_ID).First(&type_comment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบชนิดความคิดเห็น"})
		return
	}

	// 8: ค้นหา review_point ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Review_point_ID).First(&review_point); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบคะแนนรีวิว"})
		return
	}


	// 9: ค้นหา payment ด้วย id
	if tx := entity.DB().Where("id = ?", comment.Payment_ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบใบเสร็จ"})
		return
	}

	// 12: สร้าง Comment
	wv := entity.Comment{
		Comments:		   comment.Comments,      // ตั้งค่าฟิลด์ Comments
		Review_point:      review_point,          // โยงความสัมพันธ์กับ Entity Review_Point
		Payment:           payment,               // โยงความสัมพันธ์กับ Entity Payment
		Type_Com:          type_comment,          // โยงความสัมพันธ์กับ Entity Type_Comment
		Date_Now:          comment.Date_Now,      // ตั้งค่าฟิลด์ Date_Now
		Bought_now:        comment.Bought_now,    // ตั้งค่าฟิลด์ Bought_now
	}

	// 11, 12, 13: ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 14: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /comment/:id
func GetComment(c *gin.Context) {
	var commentS entity.Comment
	id := c.Param("id")
	if err := entity.DB().Preload("Review_point").Preload("Payment").Preload("Type_Com").Raw("SELECT * FROM comments WHERE id = ?", id).Find(&commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": commentS})
}

// GET /comments
func ListComments(c *gin.Context) {
	var commentS []entity.Comment
	if err := entity.DB().Preload("Review_point").Preload("Payment").Preload("Type_Com").Raw("SELECT * FROM comments").Find(&commentS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentS})
}

// DELETE /comments/:id
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM comments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบความคิดเห็น"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /comments
func UpdateComment(c *gin.Context) {

	var commentS entity.Comment
	id := c.Param("id")
	var type_com entity.Type_Comment
	var review_point entity.Review_Point
	var pay entity.Payment

	if err := c.ShouldBindJSON(&commentS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", commentS.Type_Com_ID).First(&type_com); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบชนิดความคิดเห็น"})
		return
	}
	if tx := entity.DB().Where("id = ?", commentS.Review_point_ID).First(&review_point); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบคะแนนรีวิว"})
		return
	}
	if tx := entity.DB().Where("id = ?", commentS.Payment_ID).First(&pay); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบใบเสร็จ"})
		return
	}
	dc := entity.Comment{           
		Comments:		    commentS.Comments,     
		Review_point:	    review_point,               
		Payment:	    	pay,               
		Type_Com:	    	type_com,               
		Date_Now:      		commentS.Date_Now,
		Bought_now:	        commentS.Bought_now,  
	}
	
	if _, err := govalidator.ValidateStruct(dc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Where("id = ?", id).Updates(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentS})


}
