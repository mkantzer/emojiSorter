package emojihdl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkantzer/emojiSorter/internal/core/ports"
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
)

type HTTPHandler struct {
	emojiService ports.EmojiService
}

func NewHTTPHandler(emojiService ports.EmojiService) *HTTPHandler {
	return &HTTPHandler{
		emojiService: emojiService,
	}
}

func (hdl *HTTPHandler) Health(c *gin.Context) {
	c.String(200, "healthy")
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	emoji, err := hdl.emojiService.Get(name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrEmojiNotFound):
			c.AbortWithError(404, apperrors.ErrEmojiNotFound)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, emoji)
}

func (hdl *HTTPHandler) GetAll(c *gin.Context) {
	emoji, err := hdl.emojiService.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrEmojiNotFound):
			c.AbortWithError(404, apperrors.ErrEmojiNotFound)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, emoji)
}
