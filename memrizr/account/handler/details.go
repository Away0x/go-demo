package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Details",
	})
}
