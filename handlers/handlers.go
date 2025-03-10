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

	r.POST("/login", controlers.Login)
	r.POST("/send/secret/code", controlers.SendSecretCode) 
	r.POST("/update/password", controlers.UpdateAdminPassword)
	

	r.POST("/add/admin", createadmin.AdminRegistration)
	r.POST("/update/admin", createadmin.UpdateAdmin)
	r.DELETE("/delete/admin", controlers.DeleteAdmin)
	r.GET("/get/admin", controlers.GetAdmins)


	r.POST("/add/statistic", controlers.AddStatistic)
	r.POST("/update/statistic", controlers.UpdateStatistic)
	r.GET("/get/statistic", controlers.GetStatistics)


	r.POST("/add/statistic/for/project", controlers.AddStatisticForProject) 
	r.POST("/update/project/statistic", controlers.UpdateProjectStatistic) 
	r.GET("/get/project/statistic", controlers.GetStatisticsForProject) 


	r.POST("/add/patient/story", controlers.AddPatientStory) 
	r.DELETE("/delete/patient/story", controlers.DeletePatientStory) 
	r.GET("/get/patient/story", controlers.GetPatientStories) 



	r.POST("/add/partner", controlers.AddPartner)
	r.DELETE("/delete/partner", controlers.DeletePartners) 
	r.GET("/get/partner", controlers.GetPartners) 


	r.POST("/add/team", controlers.AddTeamMembers) 
	r.DELETE("/delete/team", controlers.DeleteTeam)
	r.GET("/get/team", controlers.GetTeam)


	r.POST("/add/services", controlers.AddServices) 
	r.DELETE("/delete/services", controlers.DeleteServices) 
	r.GET("/get/services", controlers.GetServices)

	
	r.POST("/add/program", controlers.AddProgram)
	r.DELETE("/delete/program", controlers.DeleteProgram)
	r.GET("/get/program", controlers.GetPrograms)
	

	r.POST("/update/center/statistic", controlers.UpdateCenterStatistic) 
	r.GET("/get/center", controlers.GetCenterNumbers) 
	
	
	
	r.Run(":2020")

}


