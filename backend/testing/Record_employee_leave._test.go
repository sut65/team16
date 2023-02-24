package test

import (
	"testing"
	"time"


	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNumberEmout_notnull(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check Number_Em check", func(t *testing.T) {
		dis := entity.Record_employee_leave{
			Time_OUT: time.Now(),
			Number_Em: "", //ผิด
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("โปรดใส่เบอร์โทร"))
	})
}

func TestNumberEmout_morethen10(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check Number_Em check", func(t *testing.T) {
		dis := entity.Record_employee_leave{
			Time_OUT: time.Now(),
			Number_Em: "12345678901", //ผิด
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("เบอร์โทรควรมีน้อยกว่า 11 ตัว"))
	})
}

func TestNumberEmout_lessthen10(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check Number_Em check", func(t *testing.T) {
		dis := entity.Record_employee_leave{
			Time_OUT: time.Now(),
			Number_Em: "123456789", //ผิด
		}
	ok, err := govalidator.ValidateStruct(dis)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("เบอร์โทรควรมีมากกว่า 9 ตัว"))
	})
}

