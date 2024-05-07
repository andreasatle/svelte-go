package routes

import (
	"example.com/svelte-go/handlers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/ping", handlers.Ping)
	api.GET("/users", handlers.GetUsers)
	r.GET("/", handlers.Frontend)
}