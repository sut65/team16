package test

import (
	"testing"
	"time"
	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestAmountNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Amount cannot be blank", func(t *testing.T) {
		sep := entity.Separation{
			Date_Out: time.Now(),
			// Amount:   ,
			Status:   "+",
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
		sep := entity.Separation{
			Date_Out: time.Now(),
			Amount:   -2,
			Status:   "+",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกจำนวนเต็มบวกเท่านั้น"))
		// g.Expect(err.Error()).To(Equal("test"))
	})

}

func TestStatusNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Status cannot be blank", func(t *testing.T) {
		sep := entity.Separation{
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

	t.Run("Status must be Plus or Minus", func(t *testing.T) {
		sep := entity.Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "O",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอก + หรือ - เท่านั้น"))
	})
	
}

func TestDate_OutCannotBePast(t *testing.T) {
	g := NewGomegaWithT(t)
	// today := time.Now();
	// tomorrow := today.Add(24 * time.Hour);
	yesterday := time.Now().AddDate(0, 0, -1);

	t.Run("Date Out cannot be past", func(t *testing.T) {
		sep := entity.Separation{
			Date_Out: yesterday,
			Amount:   2,
			Status:   "+",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นอดีต"))
	})

}





