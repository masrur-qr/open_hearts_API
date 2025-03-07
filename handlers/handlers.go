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

	r.POST("/add/Admin",createadmin.AdminRegistration)
	r.POST("/add/statistic",controlers.AddStatistic)
	r.POST("/add/pationt/story",controlers.AddPatientStory)

	r.POST("/add/partner",controlers.AddPartner)
	r.POST("/add/statistic/for/project",controlers.Add_statistic_for_project)
	r.POST("/add/team",controlers.AddTeamMambers)
	r.POST("/add/servisec",controlers.AddServices)
	r.POST("/add/program",controlers.AddProgram)
	
	
	
	r.POST("/update/Admin",createadmin.UpdateAdmin)
	r.POST("/update/Statistic",controlers.UpdateStatistic)
	r.POST("/update/Project/Sataistic",controlers.UpdateProjectStatistic)
	r.POST("/update/Center/Sataistic",controlers.UpdateCenterStatistic)




	r.Run(":2020")
}


