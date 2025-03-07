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

	r.POST("/addAdmin",createadmin.AdminRegistration)
	r.POST("/add_statistic",controlers.AddStatistic)
	r.POST("/add_pationt_story",controlers.AddPatientStory)
	r.POST("/add_partner",controlers.AddPartner)
	r.POST("/add_statistic_for_project",controlers.Add_statistic_for_project)
	r.POST("/add_Team",controlers.AddTeamMambers)
	r.POST("/add_servisec",controlers.AddServices)
	r.POST("/add_program",controlers.AddProgram)
	
	
	
	r.POST("/updateAdmin",createadmin.UpdateAdmin)
	r.POST("/updateStatistic",controlers.UpdateStatistic)
	r.POST("/updateProjectSataistic",controlers.UpdateProjectStatistic)
	r.POST("/updateCenterSataistic",controlers.UpdateCenterStatistic)

	


	r.Run(":2020")
}


