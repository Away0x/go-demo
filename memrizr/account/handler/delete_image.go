package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Image",
	})
}
