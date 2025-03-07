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
	r.POST("/login",controlers.Login)
	r.POST("/updateAdmin",createadmin.UpdateAdmin)
	r.POST("/addAdmin",createadmin.AdminRegistration)
	r.POST("/add_statistic",controlers.AddStatistic)



	r.Run(":2020")
}


