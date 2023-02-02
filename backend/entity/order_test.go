package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestQuantityNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Quantity can not Blank", func(t *testing.T) {
		order := Order{
			Prices: 1,
		}

		ok, err := govalidator.ValidateStruct(order)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("ระบุจำนวน"))
	})
}
func TestQuantityInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Quantity must be in range (1-100)", func(t *testing.T) {
		order := Order{
			Quantity: 101,
			Prices: 1,
		}

		ok, err := govalidator.ValidateStruct(order)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("จำนวนอยู่ในช่วง (1-100)"))
	})
}

func TestPricesInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Price must be in range (1-1000000)", func(t *testing.T) {
		order := Order{
			Quantity: 10,
			Prices: -5,
		}

		ok, err := govalidator.ValidateStruct(order)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("ราคาอยู่ในช่วง (1-1000000)"))
	})
}
