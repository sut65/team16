package entity

import (
	"time"

	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Car_Model			string
	Registation_Number	string
	Delivery  []Delivery `gorm:"foreignKey:Car_ID"`
}

type Delivery struct {
	gorm.Model
	Location string			`json:"location" valid:"required~กรุณากรอกสถานที่"`
	Customer_name string	`json:"customer_name" valid:"required~กรุณากรอกชื่อลูกค้า"`
	Delivery_date time.Time

	Car_ID	*uint
	Car	Car
	Employee_ID *uint
	Employee    Employee
	Payment_ID		*uint
	Payment	Payment
}

func TestNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check location not blank", func(t *testing.T) {
		deli := Delivery{
			Location: "",
			Customer_name:"บอย ประกรลัม",
		}
	ok, err := govalidator.ValidateStruct(deli)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกสถานที่"))
	})

	t.Run("check customer_name not blank", func(t *testing.T) {
		deli := Delivery{
			Location: "146 ม.1 ต.สำราญ อ.บ้านผือ จ.อุดรธานี",
			Customer_name:"",
		}
	ok, err := govalidator.ValidateStruct(deli)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกชื่อลูกค้า"))
	})
}
