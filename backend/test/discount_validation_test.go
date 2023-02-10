package test

import (
	"testing"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestDiscount_PriceInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check discount_price in range", func(t *testing.T) {
		dis := entity.Discount{
			Discount_Price: -1,
			Discount_s: time.Now(),
			Discount_e: time.Now(),
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกราคาที่ลดอยู่ในช่วง 1-1000"))
	})
}

func TestDiscount_sNotPast(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	yesterday := time.Now().AddDate(0, 0, -1);
	t.Run("check discount_s not past", func(t *testing.T) {
		dis := entity.Discount{
			Discount_Price: 1,
			Discount_s: yesterday,
			Discount_e: time.Now(),
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("วันที่เริ่มลดราคาต้องไม่เป็นวันที่ผ่านมาแล้ว"))
	})
}

func TestDiscount_eNotPast(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	yesterday := time.Now().AddDate(0, 0, -1);
	t.Run("check discount_e not past", func(t *testing.T) {
		dis := entity.Discount{
			Discount_Price: 1,
			Discount_s: time.Now(),
			Discount_e: yesterday,
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("วันที่สิ้นสุดการลดราคาต้องไม่เป็นวันที่ผ่านมาแล้ว"))
	})
}






