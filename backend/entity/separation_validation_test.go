package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestAmountNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Amount cannot be blank", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("จำนวนต้องไม่เป็นค่าว่าง"))
	})
	// sep := Separation{
	// 	Date_Out: time.Now(),
	// 	Amount:   5,
	// 	Status:   "+",
	// }

	// // ตรวจสอบด้วย govalidator
	// ok, err := govalidator.ValidateStruct(sep)

	// // ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	// g.Expect(ok).ToNot(BeTrue())

	// // err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	// g.Expect(err).ToNot(BeNil())

	// // err.Error ต้องมี error message แสดงออกมา
	// g.Expect(err.Error()).To(Equal("Amount cannot be blank"))
}

func TestStatusNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Status cannot be blank", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("สถานะต้องไม่เป็นค่าว่าง"))
	})
	
}

func TestStatusMustBePlusOrNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Status must be positive", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอก + หรือ - เท่านั้น"))
	})
	
}


