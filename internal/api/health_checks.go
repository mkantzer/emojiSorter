package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloServer(c *gin.Context) {
	c.String(http.StatusOK, "Hello %s!\n", "World")
}

func HealthCheck(c *gin.Context) {
	c.String(
		http.StatusOK,
		"This seems fine\n",
	)
}

func UnhealthCheck(c *gin.Context) {
	c.String(
		http.StatusInternalServerError,
		"This seems Not Fine\n",
	)
}
