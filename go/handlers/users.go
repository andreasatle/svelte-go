package handlers

import (
	"net/http"

	"example.com/svelte-go/database"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []database.User
	if err := database.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)}
