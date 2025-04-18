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
	controlers.AddStatistic()
	controlers.AddStatisticForProject()

	r.POST("/login", controlers.Login)
	r.POST("/send/secret/code", controlers.SendSecretCode) 
	r.POST("/update/password", controlers.UpdateAdminPassword)
	

	r.POST("/add/admin", createadmin.AdminRegistration)
	r.POST("/update/admin", createadmin.UpdateAdmin)
	r.DELETE("/delete/admin", controlers.DeleteAdmin)
	r.GET("/get/admin", controlers.GetAdmins)
	r.GET("/get/one/admin", controlers.GetAdmin)


	r.POST("/update/statistic", controlers.UpdateStatistic)
	r.GET("/get/statistic", controlers.GetStatistics)


	r.POST("/update/project/statistic", controlers.UpdateProjectStatistic) 
	r.GET("/get/project/statistic", controlers.GetStatisticsForProject) 


	r.POST("/add/patient/story", controlers.AddPatientStory) 
	r.DELETE("/delete/patient/story", controlers.DeletePatientStory) 
	r.GET("/get/patient/story", controlers.GetPatientStories) 
	r.GET("/get/one/patient", controlers.GetOnePatient) 



	r.POST("/add/partner", controlers.AddPartner)
	r.DELETE("/delete/partner", controlers.DeletePartners) 
	r.GET("/get/partners", controlers.GetPartners) 


	r.POST("/add/team", controlers.AddTeamMembers) 
	r.DELETE("/delete/team", controlers.DeleteTeam)
	r.GET("/get/team", controlers.GetTeam)


	r.POST("/add/services", controlers.AddServices) 
	r.DELETE("/delete/services", controlers.DeleteServices) 
	r.GET("/get/services", controlers.GetServices)

	
	r.POST("/add/program", controlers.AddProgram)
	r.DELETE("/delete/program", controlers.DeleteProgram)
	r.GET("/get/programs", controlers.GetPrograms)
	

	r.POST("/update/center/statistic", controlers.UpdateCenterStatistic) 
	r.GET("/get/center", controlers.GetCenterNumbers) 

	
	
	r.GET("/read/photo",controlers.ReadFile)
	r.Run(":2020")

}


