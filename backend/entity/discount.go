package entity

import (
	_"time"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Name      string
}