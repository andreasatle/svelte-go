package main

import (
	"log"
	"net/http"

	"example.com/svelte-go/config"
	"example.com/svelte-go/database"
	"example.com/svelte-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	database.ConnectDB(config.DBConfig)
	r := gin.Default()
	r.StaticFS("/pages", http.Dir("../svelte/build"))
	routes.LoadRoutes(r)

	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatalf("Certification error: %v", err)
	}
}
