package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Server) testLog(c *gin.Context) {
	a.Deps.Logger.Info("hey look at me!")
	c.String(http.StatusOK, "Hewdy %s", "bruh")
}

// func (a *Server)
