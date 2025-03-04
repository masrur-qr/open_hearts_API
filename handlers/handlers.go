package handlers

import (
	controlers "docs/app"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.Use(controlers.Cors)
}


