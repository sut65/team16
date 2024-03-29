package test

import (
	"testing"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStock(t *testing.T) {
	g := NewGomegaWithT(t)

	st := entity.Stock{
		Name:     "Me",
		Amount:   100,
		Price:    60.00,
		DateTime: time.Now(),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(st)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())

}

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	st := entity.Stock{
		Name:     "",
		Amount:   100,
		Price:    60,
		DateTime: time.Now(),
	}

	ok, err := govalidator.ValidateStruct(st)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ชื่อสินค้าห้ามเป็นค่าว่าง"))
}

// Amount
func TestAmount(t *testing.T) {
	g := NewGomegaWithT(t)
	fixture := []int{
		0, 1001,
	}

	for _, am := range fixture {
		st := entity.Stock{
			Name:     "Me",
			Amount:   am,
			Price:    60,
			DateTime: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(st)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("จำนวนสินค้าต้องอยู่ในช่วง 1 - 1000"))
	}
}

// Price
func TestPrice(t *testing.T) {
	g := NewGomegaWithT(t)
	fixture := []float64{
		0, 1001,
	}

	for _, pr := range fixture {
		st := entity.Stock{
			Name:     "Me",
			Amount:   100,
			Price:    pr,
			DateTime: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(st)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ราคาสินค้าต้องอยู่ในช่วง 1 - 1000 บาท"))
	}
}

// DateTime Past-Future
func TestDateTimeMustNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	st := entity.Stock{
		Name:     "Me",
		Amount:   100,
		Price:    60.00,
		DateTime: time.Now().Add(-24 * time.Hour),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(st)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).NotTo(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("เวลาต้องไม่เป็นอดีต"))
}

func TestDateTimeMustNotBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	st := entity.Stock{
		Name:     "Me",
		Amount:   100,
		Price:    60.00,
		DateTime: time.Now().Add(24 * time.Hour), //ผิด
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(st)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).NotTo(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("เวลาต้องไม่เป็นอนาคต"))
}
