package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStock(t *testing.T) {
	g := NewGomegaWithT(t)

	st := Stock{
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


func TestDateTimeMustNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	st := Stock{
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
	g.Expect(err.Error()).To(Equal("DateTime must not be in the past"))
}

func TestDateTimeMushNotBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	st := Stock{
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
	g.Expect(err.Error()).To(Equal("DateTime must not be in the future"))
}
