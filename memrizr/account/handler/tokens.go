package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Tokens",
	})
}
