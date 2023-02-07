package entity

import (
	"time"

	"gorm.io/gorm"
)


type Review_Point struct {
	gorm.Model
	Point        int
	Point_Name   string
	Comment      []Comment `gorm:"foreignKey:Review_point_ID"`
}

type Type_Comment struct {
	gorm.Model
	Type_Com_Name       string
	Comment             []Comment `gorm:"foreignKey:Type_Com_ID"`
}

type Comment struct {
	gorm.Model
	Comments    string			`valid:"required~ความคิดเห็นต้องไม่เป็นค่าว่าง"`

	Review_point_ID  *uint
	Review_point      Review_Point	`valid:"-"`

	Payment_ID *uint
	Payment     Payment			`valid:"-"`

	Type_Com_ID *uint
	Type_Com    Type_Comment	`valid:"-"`

	Date_Now    time.Time		`valid:"required~ความคิดเห็นต้องไม่เป็นค่าว่าง, Past~วันที่ต้องไม่เป็นอดีต, Future~วันที่ต้องไม่เป็นอนาคต"`
	Bought_now  int				`valid:"required~จำนวนต้องไม่เป็นค่าว่าง, range(0|9223372036854775807)~กรุณากรอกจำนวนเต็มบวกเท่านั้น"`
}
