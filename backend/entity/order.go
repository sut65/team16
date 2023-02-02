package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity         int		`valid:"required~ระบุจำนวน,range(1|100)~จำนวนอยู่ในช่วง (1-100)"` 
	Prices           float64	`valid:"required~ระบุราคา,range(1|1000000)~ราคาอยู่ในช่วง (1-1000000)"` 
	Shelving_ID      *uint
	Shelving         Shelving	
	Shopping_Cart_ID *uint		
	Shopping_Cart    Shopping_Cart
}

type Status struct {
	gorm.Model
	Status string

	Shopping_Cart []Shopping_Cart `gorm:"foreignKey:Status_ID"`
}

func TestOrderNotZero(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Quantity can not Zero", func(t *testing.T) {
		order := Order{
			Quantity: -1,
			Prices: 1,
		}

		ok, err := govalidator.ValidateStruct(order)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("จำนวนสินค้าต้องเป็นจำนวนเต็ม"))
	})

	t.Run("Prices can not Zero", func(t *testing.T) {
		order := Order{
			Quantity: 1,
			Prices: -1,
		}

		ok, err := govalidator.ValidateStruct(order)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Prices in range (0,10000)"))
	})
}

