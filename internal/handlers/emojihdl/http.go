package emojihdl

import (
	"github.com/gin-gonic/gin"
	"github.com/mkantzer/emojiSorter/internal/core/ports"
)

type HTTPHandler struct {
	emojiService ports.EmojiService
}

func NewHTTPHandler(emojiService ports.EmojiService) *HTTPHandler {
	return &HTTPHandler{
		emojiService: emojiService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	emoji, err := hdl.emojiService.Get(c.Param("name"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, emoji)
}
