package test

import (
	"testing"
	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)


func TestMNameNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check Member Name can  not blank", func(t *testing.T) {
		member := entity.Member{
			Mem_Name: "",
			Mem_Age:  15,
			Mem_Tel:  "0930963238",
		}

		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกชื่อ - นามสกุล"))
	})
}
func TestAgeMT15(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check Member Age must be more than 15 years old", func(t *testing.T) {
		member :=  entity.Member{
			Mem_Name: "keng",
			Mem_Age:  10,
			Mem_Tel:  "0930963238",
		}

		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("โปรดระบุอายุที่มากกว่า 15 ปีขึ้นไป"))
	})
}
func TestTelNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check Member Telephone number cannot be blank", func(t *testing.T) {
		member :=  entity.Member{
			Mem_Name: "Keng",
			Mem_Age:  16,
			Mem_Tel:  "",
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกเบอร์มือถือ"))
	})
}