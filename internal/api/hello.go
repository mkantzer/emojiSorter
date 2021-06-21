package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloServer(c *gin.Context) {
	c.String(http.StatusOK, "Hello %s!\n", "World")
}
