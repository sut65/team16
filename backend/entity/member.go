package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Gender_Name string
	Member      []Member `gorm:"foreignkey:Gender_ID"`
}

type Level struct {
	gorm.Model
	Level_Name    string
	Level_Pay     int
	Level_Benefit string
	Member        []Member `gorm:"foreignkey:Level_ID"`
}

type Member struct {
	gorm.Model
	Mem_Name string `json:"Mem_Name" valid:"required~กรุณากรอกชื่อ - นามสกุล"`
	Mem_Age  int `json:"Mem_Age" valid:"range(15|100)~โปรดระบุอายุที่มากกว่า 15 ปีขึ้นไป"`
	Mem_Tel  string `gorm:"uniqueIndex" json:"Mem_Tel" valid:"required~กรุณากรอกเบอร์มือถือ"`

	Gender_ID   *uint`valid:"-"`
	Gender      Gender`valid:"-"`
	Level_ID    *uint`valid:"-"`
	Level       Level`valid:"-"`
	Employee_ID *uint`valid:"-"`
	Employee    Employee`valid:"-"`

	Shopping_Cart []Shopping_Cart `gorm:"foreignkey:Member_ID"`
}

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Member Name can  not blank", func(t *testing.T) {
		member := Member{
			Mem_Name: "",
			Mem_Age:  15,
			Mem_Tel:  "0930963238",
		}

		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกชื่อ - นามสกุล"))
	})
}
func TestAgeMT15(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("check Member Age must be more than 15 years old", func(t *testing.T) {
		member := Member{
			Mem_Name: "keng",
			Mem_Age:  10,
			Mem_Tel:  "0930963238",
		}

		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("โปรดระบุอายุที่มากกว่า 15 ปีขึ้นไป"))
	})
}
func TestTelNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("check Member Telephone number cannot be blank", func(t *testing.T) {
		member := Member{
			Mem_Name: "Keng",
			Mem_Age:  16,
			Mem_Tel:  "",
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกเบอร์มือถือ"))
	})
}
