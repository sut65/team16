package main

import (
	"github.com/Team16/farm_mart/controller"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Employee Routes
	r.GET("/emloyees", controller.ListEmployees)
	r.GET("/emloyee/:id", controller.GetEmployee)
	r.POST("/emloyees", controller.CreateEmployee)
	r.PATCH("/emloyees", controller.UpdateEmployee)
	r.DELETE("/emloyees/:id", controller.DeleteEmployee)

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
	r.GET("/Shelving", controller.ListShelvings)
	r.GET("/Shelving/:id", controller.GetShelving)
	r.POST("/Shelving", controller.CreateShelving)
	r.PATCH("/Shelving", controller.UpdateShelving)
	r.DELETE("/Shelving/:id", controller.DeleteShelving)

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

	// Run the server
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
