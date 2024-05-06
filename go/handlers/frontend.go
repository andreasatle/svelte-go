package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Frontend(c *gin.Context) {
	c.Redirect(http.StatusFound, "/pages")
}