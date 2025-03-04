package handlers

import (
	"docs/app/controlers"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.Use(controlers.Cors)
}


