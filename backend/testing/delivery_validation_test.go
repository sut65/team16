package test

import (
	"testing"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestCustomer_nameNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check location not blank", func(t *testing.T) {
		deli := entity.Delivery{
			Location: "",
			Customer_name:"บอย ประกรลัม",
			Delivery_date: time.Now(),
		}
	ok, err := govalidator.ValidateStruct(deli)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกสถานที่"))
	})

}

func TestLocationNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check customer_name not blank", func(t *testing.T) {
		deli := entity.Delivery{
			Location: "146 ม.1 ต.สำราญ อ.บ้านผือ จ.อุดรธานี",
			Customer_name:"",
			Delivery_date: time.Now(),
		}
	ok, err := govalidator.ValidateStruct(deli)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกชื่อลูกค้า"))
	})
}

func TestDelivery_dateNotFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	today := time.Now();
	tomorrow := today.Add(24 * time.Hour);
	t.Run("delivery_date cannot be future", func(t *testing.T) {
		deli := entity.Delivery{
			Location: "146 ม.1 ต.สำราญ อ.บ้านผือ จ.อุดรธานี",
			Customer_name:"บอย ประกรลัม",
			Delivery_date: tomorrow,
		}
	ok, err := govalidator.ValidateStruct(deli)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("วันที่ส่งสินค้าต้องไม่เป็นวันที่ในอนาคต"))
	})
}






