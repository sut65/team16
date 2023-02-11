package test

import (
	"testing"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestShelving(t *testing.T) {
	g := NewGomegaWithT(t)

	sv := entity.Shelving{
		Number:    15,
		Cost:      60.00,
		Date_Time: time.Now(),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(sv)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())

}
func TestNumberOf(t *testing.T) {
	g := NewGomegaWithT(t)
	fixture := []int{
		0, 21,
	}

	for _, sh := range fixture {
		sh := entity.Shelving{
			Number:    sh,
			Cost:      20,
			Date_Time: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(sh)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Number must be in the range 1-20"))
	}

}

func TestCost(t *testing.T) {
	g := NewGomegaWithT(t)
	fixture := []float64{
		0, 1001,
	}

	for _, ct := range fixture {
		sv := entity.Shelving{

			Number:    20,
			Cost:      ct,
			Date_Time: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(sv)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Cost must be in the range 1-1000"))
	}
}

// DateTime Past-Future
func TestDateTimeNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	sv := entity.Shelving{

		Number:    20,
		Cost:      30,
		Date_Time: time.Now().Add(-24 * time.Hour),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(sv)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).NotTo(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("DateTime must not be in the past"))
}

func TestDateTimeNotBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	sv := entity.Shelving{

		Number:    20,
		Cost:      35,
		Date_Time: time.Now().Add(24 * time.Hour), //ผิด
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(sv)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).NotTo(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("DateTime must not be in the future"))
}
