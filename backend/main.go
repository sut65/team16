package main

import (
	"github.com/Team16/farm_mart/controller"
	"github.com/Team16/farm_mart/entity"
	"github.com/Team16/farm_mart/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	router := gin.Default()
	router.Use(CORSMiddleware())

	r := router.Group("/")

	{
		r.Use(middlewares.Authorizes())
		{

			// Employee Routes
			r.GET("/employees", controller.ListEmployees)
			r.GET("/employee/:id", controller.GetEmployee)
			r.POST("/employees", controller.CreateEmployee)
			r.PATCH("/employees", controller.UpdateEmployee)
			r.DELETE("/employees/:id", controller.DeleteEmployee)

			//Kind Routes
			r.GET("/kinds", controller.ListKinds)
			r.GET("/kind/:id", controller.GetKind)
			r.POST("/kinds", controller.CreateKind)
			r.PATCH("/kinds", controller.UpdateKind)
			r.DELETE("/kinds/:id", controller.DeleteKind)

			//Stock Routes
			r.GET("/stocks", controller.ListStocks)
			r.GET("/stock/:id", controller.GetStock)
			r.POST("/stocks", controller.CreateStock)
			r.PATCH("/stocks/:id", controller.UpdateStock)
			r.DELETE("/stocks/:id", controller.DeleteStock)

			//Labels Routes
			r.GET("/labels", controller.ListLabels)
			r.GET("/label/:id", controller.GetLabel)
			r.POST("/labels", controller.CreateLabel)
			r.PATCH("/labels", controller.UpdateLabel)
			r.DELETE("/labels/:id", controller.DeleteLabel)

			//Shelving Routes
			r.GET("/shelvings", controller.ListShelvings)
			r.GET("/shelvings/:id", controller.ListLabel)
			r.GET("/shelving/:id", controller.GetShelving)
			r.POST("/shelvings", controller.CreateShelving)
			r.PATCH("/shelves/:id", controller.UpdateShelving)
			r.PATCH("/shelf/:id", controller.UpdateNumber)
			r.DELETE("/shelvings/:id", controller.DeleteShelving)

			//Member Routes
			r.GET("/members", controller.ListMember)
			r.GET("/member/:id", controller.GetMember)
			r.POST("/members", controller.CreateMember)
			r.PATCH("/members/:id", controller.UpdateMember)
			r.DELETE("/members/:id", controller.DeleteMember)

			//Gender Routes
			r.GET("/genders", controller.ListGender)
			r.GET("/gender/:id", controller.GetGender)
			r.POST("/genders", controller.CreateGender)
			r.PATCH("/genders/:id", controller.UpdateGender)
			r.DELETE("/genders/:id", controller.DeleteGender)

			//Level Routes
			r.GET("/levels", controller.ListLevel)
			r.GET("/level/:id", controller.GetLevel)
			r.POST("/levels", controller.CreateLevel)
			r.PATCH("/levels/:id", controller.UpdateLevel)
			r.DELETE("/levels/:id", controller.DeleteLevel)

			//Leave Routes
			r.GET("/leaves", controller.ListLeave)
			r.GET("/leave/:id", controller.GetLeave)
			r.POST("/leaves", controller.CreateLeave)
			r.PATCH("/leaves/:id", controller.UpdateLeave)
			r.DELETE("/leaves/:id", controller.DeleteLeave)

			//Section Routes
			r.GET("/sections", controller.ListSection)
			r.GET("/section/:id", controller.GetSection)
			r.POST("/sections", controller.CreateSection)
			r.PATCH("/sections/:id", controller.UpdateSection)
			r.DELETE("/sections/:id", controller.DeleteSection)

			//L_Type Routes
			r.GET("/l_types", controller.ListL_Type)
			r.GET("/l_type/:id", controller.GetL_Type)
			r.POST("/l_types", controller.CreateL_Type)
			r.PATCH("/l_types/:id", controller.UpdateL_Type)
			r.DELETE("/l_types/:id", controller.DeleteL_Type)

			//Employee_attendence Routes
			r.GET("/employee_attendances", controller.ListEmployee_attendance)
			r.GET("/employee_attendances/:id", controller.GetEmployee_attendance)
			r.POST("/employee_attendances", controller.CreateEmployee_attendance)
			r.PATCH("/employee_attendances/:id", controller.UpdateEmployee_attendance)
			r.DELETE("/employee_attendances/:id", controller.DeleteEmployee_attendance)

			//Record_employee_leave Routes
			r.GET("/record_employee_leaves", controller.ListRecord_employee_leave)
			r.GET("/record_employee_leaves/:id", controller.GetRecord_employee_leave)
			r.POST("/record_employee_leaves", controller.CreateRecord_employee_leave)
			r.PATCH("/record_employee_leaves/:id", controller.UpdateRecord_employee_leave)
			r.DELETE("/record_employee_leaves/:id", controller.DeleteRecord_employee_leave)

			//list components employee attendances
			r.GET("/listduty", controller.Listduty)
			r.GET("/listworking_time", controller.Listworking_time)
			r.GET("/listovertime", controller.Listovertime)

			//Discount Routes
			r.GET("/discounts", controller.ListDiscount)
			r.GET("/discount/:id", controller.GetDiscount)
			r.POST("/discounts/:shelvingID", controller.CreateDiscount)
			r.PATCH("/discount/:id", controller.UpdateDiscount)
			r.DELETE("/discount/:id", controller.DeleteDiscount)

			//Discouting price Routes
			r.PATCH("/discounting/:id", controller.DiscountingShelving)

			//Reset price Routes
			r.PATCH("/reset_cost/:shelvingID/:cost", controller.ResetCost)

			//Discount_Type Routes
			r.GET("/discount_types", controller.ListDiscount_Type)

			//Separartion Routes
			r.GET("/separations", controller.ListSeparations)
			r.GET("/separation/:id", controller.GetSeparation)
			r.POST("/separations", controller.CreateSeparation)
			r.PATCH("/separation/:id", controller.UpdateSeparation)
			r.DELETE("/separation/:id", controller.DeleteSeparation)

			//Storage Routes
			r.GET("/storages", controller.ListStorages)
			r.GET("/storage/:id", controller.GetStorage)
			r.POST("/storages", controller.CreateStorage)
			r.PATCH("/storages", controller.UpdateStorage)
			r.DELETE("/Storages/:id", controller.DeleteStorage)

			// //Comment Routes
			// r.GET("/comments", controller.ListComments)
			// r.GET("/comment/:id", controller.GetComment)
			// r.POST("/comments", controller.CreateComment)
			// r.PATCH("/comment/:id", controller.UpdateComment)
			// r.DELETE("/comment/:id", controller.DeleteComment)

			// //Type comment Routes
			// r.GET("/type_comments", controller.ListType_comments)
			// r.GET("/type_comment/:id", controller.GetType_comment)
			// r.POST("/type_comments", controller.CreateType_comment)
			// r.PATCH("/type_comment/:id", controller.UpdateType_comment)
			// r.DELETE("/type_comment/:id", controller.DeleteType_comment)

			// //Review point Routes
			// r.GET("/review_points", controller.ListReview_points)
			// r.GET("/review_point/:id", controller.GetReview_point)
			// r.POST("/review_points", controller.CreateReview_point)
			// r.PATCH("/review_point/:id", controller.UpdateReview_point)
			// r.DELETE("/review_point/:id", controller.DeleteReview_point)

			//Reason Routes
			r.GET("/reasons", controller.ListReasons)
			r.GET("/reason/:id", controller.GetReason)
			r.POST("/reasons", controller.CreateReason)
			r.PATCH("/reason/:id", controller.UpdateVideo)
			r.DELETE("/reason/:id", controller.DeleteReason)

			//Order Routes
			r.GET("/orders", controller.ListOrder)
			r.GET("/order/:id", controller.GetOrder)
			r.GET("/shelv", controller.ListShelv)
			r.GET("/ordercart/:id", controller.ListOrderCart)
			r.GET("/ordersum/:id", controller.ListOrdersum)
			r.POST("/orders", controller.CreateOrder)
			r.PATCH("/order/:id", controller.UpdateOrder)
			r.PATCH("/UpdateQuantity/:id", controller.UpdateQuantity)
			r.DELETE("/order/:id", controller.DeleteOrder)

			//payment Routes
			r.GET("/payments", controller.ListPayment)
			r.GET("/payment/:id", controller.GetPayment)
			r.POST("/payments", controller.CreatePayment)
			r.PATCH("/payment/:id", controller.UpdatePayment)
			r.DELETE("/payment/:id", controller.DeletePayment)

			//payment_methods Routes
			r.GET("/payment_methods", controller.Listpayment_method)
			r.GET("/payment_method/:id", controller.GetPayment_method)
			r.POST("/payment_methods", controller.CreatePayment_method)
			r.PATCH("/payment_method/:id", controller.UpdatePayment_method)
			r.DELETE("/payment_method/:id", controller.DeletePayment_method)

			//Status Routes
			r.GET("/statuses", controller.ListStatus)
			r.GET("/status/:id", controller.GetStatus)
			r.POST("/statuses", controller.CreateStatus)
			r.PATCH("/status/:id", controller.UpdateStatus)
			r.DELETE("/status/:id", controller.DeleteStatus)

			//Status Routes
			r.GET("/carts", controller.ListCart)
			r.GET("/unpaids", controller.ListUnpaid)
			r.GET("/paids", controller.Listpaid)
			r.GET("/cart/:id", controller.GetCart)
			r.POST("/carts", controller.CreateCart)
			r.PATCH("/cart/:id", controller.UpdateCart)
			r.DELETE("/cart/:id", controller.DeleteCart)

			//Car Routes
			r.GET("/cars", controller.ListCar)

			//Delivery Routes
			r.GET("/deliveries", controller.ListDelivery)
			r.GET("/delivery/:id", controller.GetDelivery)
			r.POST("/deliveries", controller.CreateDelivery)
			r.PATCH("/delivery/:id", controller.UpdateDelivery)
			r.DELETE("/delivery/:id", controller.DeleteDelivery)
		}
	}
	//Comment Routes
	router.GET("/comments", controller.ListComments)
	router.GET("/comment/:id", controller.GetComment)
	router.POST("/comments", controller.CreateComment)
	router.PATCH("/comment/:id", controller.UpdateComment)
	router.DELETE("/comment/:id", controller.DeleteComment)

	//Type comment Routes
	router.GET("/type_comments", controller.ListType_comments)
	router.GET("/type_comment/:id", controller.GetType_comment)
	router.POST("/type_comments", controller.CreateType_comment)
	router.PATCH("/type_comment/:id", controller.UpdateType_comment)
	router.DELETE("/type_comment/:id", controller.DeleteType_comment)

	//Review point Routes
	router.GET("/review_points", controller.ListReview_points)
	router.GET("/review_point/:id", controller.GetReview_point)
	router.POST("/review_points", controller.CreateReview_point)
	router.PATCH("/review_point/:id", controller.UpdateReview_point)
	router.DELETE("/review_point/:id", controller.DeleteReview_point)

	router.POST("/login", controller.Login)
	router.POST("/Access", controller.Access)
	// Run the server
	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
