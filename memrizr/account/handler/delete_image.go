package handler

import (
	"log"
	"net/http"

	"memrizr/model"
	"memrizr/model/apperrors"

	"github.com/gin-gonic/gin"
)

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	authUser := c.MustGet("user").(*model.User)

	ctx := c.Request.Context()
	err := h.UserService.ClearProfileImage(ctx, authUser.UID)

	if err != nil {
		log.Printf("Failed to delete profile image: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
