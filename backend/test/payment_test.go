package test

import (
	"time"
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"github.com/Team16/farm_mart/entity"
)

func TestPaymentTimeNotPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Time can not Pass", func(t *testing.T) {
		payment := entity.Payment{
			Time: time.Now().AddDate(0,0,-1),
			Paytotal: 100,
			Note: "",
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("วันที่ห้ามเป็นอดีต"))
	})
}
func TestPaymentTimeNotFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Time can not Future", func(t *testing.T) {
		payment := entity.Payment{
			Time: time.Now().AddDate(0,0,1),
			Paytotal: 100,
			Note: "",
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("วันที่ห้ามเป็นอนาคต"))
	})
}

func TestPaytotalInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Price must be in range (1-1000000)", func(t *testing.T) {
		payment := entity.Payment{
			Time: time.Now(),
			Paytotal: -1,
			Note: "",
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("ยอดรวมอยู่ในช่วง (1-1000000)"))
	})
}
func TestNoteStringLength64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Note not more than 64 char", func(t *testing.T) {
		payment := entity.Payment{
			Time: time.Now(),
			Paytotal: 10000,
			Note: "12345678901234567890123456790123457890123456789012345678901234567890",
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("อักขระไม่เกิน 64 ตัว"))
	})
}