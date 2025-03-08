package handlers

import (
	"docs/app/controlers"
	"docs/app/cors"
	"docs/app/createadmin"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.Use(cors.Cors)

	createadmin.Createadmin()
	controlers.AddStatisticForCenter()

	r.POST("/login",controlers.Login)

	r.POST("/add/admin",createadmin.AdminRegistration)
	r.POST("/add/statistic",controlers.AddStatistic)
	r.POST("/add/statistic/for/project",controlers.Add_statistic_for_project)
	r.POST("/add/pationt/story",controlers.AddPatientStory)
	r.POST("/add/partner",controlers.AddPartner)
	r.POST("/add/team",controlers.AddTeamMambers)
	r.POST("/add/servisec",controlers.AddServices)
	r.POST("/add/program",controlers.AddProgram)
	
	
	
	r.POST("/update/admin",createadmin.UpdateAdmin)
	r.POST("/update/statistic",controlers.UpdateStatistic)
	r.POST("/update/project/sataistic",controlers.UpdateProjectStatistic)
	r.POST("/update/center/sataistic",controlers.UpdateCenterStatistic)

	
	r.DELETE("/delete/pationt/story",controlers.DeletePatientStory)
	r.DELETE("/delete/partner",controlers.DeletePatners)
	r.DELETE("/delete/team",controlers.DeleteTeam)
	r.DELETE("/delete/servisec",controlers.DeleteServisec)
	r.DELETE("/delete/program",controlers.DeleteProgram)




	
	r.GET("/get/statistic",controlers.GetStatistic)
	r.GET("/get/pationt/story",controlers.GetPatientStory)
	r.GET("/get/patner",controlers.GetPatners)
	r.GET("/get/center",controlers.Get_center_number)
	r.GET("/get/project/statistic",controlers.GetStatisticforproject)
	r.GET("/get/team",controlers.GetTeam)
	r.GET("/get/services",controlers.GetServices)
	r.GET("/get/program",controlers.GetProgram)


	r.Run(":2020")
}


