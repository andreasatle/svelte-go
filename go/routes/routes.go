package routes

import (
	"example.com/svelte-go/handlers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
	r.GET("/", handlers.Frontend)
}