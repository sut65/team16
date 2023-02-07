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

}

func TestAmountMustBePositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Amount must be positive", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกจำนวนเต็มบวกเท่านั้น"))
	})

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

