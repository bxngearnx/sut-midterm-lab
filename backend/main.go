package main

import (

	// "github.com/bxngearnx/team06/controller"

	"github.com/bxngearnx/sutlab/controller"
	"github.com/bxngearnx/sutlab/entity"

	"github.com/gin-gonic/gin"
)

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

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	////////////////////////////////////////////////////////////////////
	r.POST("/admin_Login", controller.Admin_Login)
	r.POST("/Student_Login", controller.Student_Login)
	r.GET("/student/:id", controller.GetStudent)
	// r.GET("/admin/:id", controller.GetAdmin)
	//////////////////////////////////////////////////////////////////
	//combobox Prefix
	r.GET("/prefix", controller.ListPrefix)
	//combobox Gender
	r.GET("/gender", controller.ListGender)
	//combobox Province
	r.GET("/province", controller.ListProvince)
	//combobox Institute
	r.GET("/institute", controller.ListInstitute)
	//combobox Branch
	r.GET("/branch", controller.ListBranch)
	//combobox Degree
	r.GET("/degree", controller.ListDegree)
	//combobox Trimester
	r.GET("/trimester", controller.ListTrimester)

	/////////////////////////////////////////////////////////////////

	//รับข้อมูลเข้าตาราง Admin
	r.POST("/create_Admin", controller.CreateAdmiin)
	//แสดงข้อมูลตาราง Admin
	r.GET("/admin_table", controller.ListAdminTable)
	// ดึงข้อมูล admin by id
	r.GET("/admin/:id", controller.ListAdminByID)
	// แก้ไขข้อมูล admin
	r.PATCH("/update_admin", controller.UpdateAdmin)
	// ลบข้อมูล admin by id
	r.DELETE("/delete_admin/:id", controller.DeleteAdminByID)

	//////////////////////////////////////////////////////////////////
	//รับข้อมูลเข้าตาราง Postponement
	r.POST("/create_Postponement", controller.CreatePostponement)
	//แสดงข้อมูลตาราง Postponement
	r.GET("/postponement_table", controller.ListPostponementTable)
	// ดึงข้อมูล Postponement by id
	r.GET("/postponement/:id", controller.ListPostponementByID)
	// แก้ไขข้อมูล Postponement
	r.PATCH("/update_postponement", controller.UpdatePostponement)
	// ลบข้อมูล Postponement by id
	r.DELETE("/delete_postponementadmin/:id", controller.DeletePostponementByID)

	//////////////////////////////////////////////////////////////////
	r.Run()

}
