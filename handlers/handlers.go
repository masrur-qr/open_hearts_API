package handlers

import (
	"docs/app/cors"
	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.Use(cors.Cors)
}


