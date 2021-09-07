package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Signup",
	})
}
