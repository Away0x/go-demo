package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Image",
	})
}
