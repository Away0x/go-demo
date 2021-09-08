package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"memrizr/model"
	"memrizr/model/apperrors"
	"memrizr/model/mocks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteImage(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// authorized middleware user
	uid, _ := uuid.NewRandom()
	ctxUser := &model.User{
		UID: uid,
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user", ctxUser)
	})

	// this handler reuqires UserService
	mockUserService := new(mocks.MockUserService)

	NewHandler(&Config{
		R:           router,
		UserService: mockUserService,
	})

	t.Run("Clear profile image error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		clearProfileImageArgs := mock.Arguments{
			mock.AnythingOfType("*context.emptyCtx"),
			ctxUser.UID,
		}

		errorResp := apperrors.NewInternal()
		mockUserService.On("ClearProfileImage", clearProfileImageArgs...).Return(errorResp)

		request, _ := http.NewRequest(http.MethodDelete, "/image", nil)
		router.ServeHTTP(rr, request)

		respBody, _ := json.Marshal(gin.H{
			"error": errorResp,
		})

		assert.Equal(t, apperrors.Status(errorResp), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertCalled(t, "ClearProfileImage", clearProfileImageArgs...)
	})

	t.Run("Success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		// authorized middleware user - overwriting for unique mock arguments
		uid, _ := uuid.NewRandom()
		ctxUser := &model.User{
			UID: uid,
		}

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", ctxUser)
		})

		// this handler reuqires UserService
		mockUserService := new(mocks.MockUserService)

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		clearProfileImageArgs := mock.Arguments{
			mock.AnythingOfType("*context.emptyCtx"),
			ctxUser.UID,
		}

		mockUserService.On("ClearProfileImage", clearProfileImageArgs...).Return(nil)

		request, _ := http.NewRequest(http.MethodDelete, "/image", nil)
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUserService.AssertCalled(t, "ClearProfileImage", clearProfileImageArgs...)
	})
}
