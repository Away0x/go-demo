package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signin used to authenticate extant user
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Signin",
	})
}
