package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamkoller/multi-tenant-nexus-manager/internal/shared/errors"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorData  `json:"error,omitempty"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	var statusCode int
	var errorData ErrorData

	if appErr, ok := err.(errors.AppError); ok {
		errorData = ErrorData{
			Code:    appErr.Code,
			Message: appErr.Message,
			Details: appErr.Details,
		}
		statusCode = getStatusCodeFromError(appErr.Code)
	} else {
		errorData = ErrorData{
			Code:    "INTERNAL_SERVER",
			Message: "Internal server error",
		}
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, Response{
		Success: false,
		Error:   &errorData,
	})
}

func getStatusCodeFromError(code string) int {
	switch code {
	case "NOT_FOUND":
		return http.StatusNotFound
	case "INVALID_INPUT":
		return http.StatusBadRequest
	case "UNAUTHORIZED":
		return http.StatusUnauthorized
	case "FORBIDDEN":
		return http.StatusForbidden
	case "CONFLICT":
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
