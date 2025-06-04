package errors

import "fmt"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code, message string) AppError {
	return AppError{
		Code:    code,
		Message: message,
	}
}

func NewAppErrorWithDetails(code, message, details string) AppError {
	return AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

var (
	ErrNotFound       = NewAppError("NOT_FOUND", "Resource not found")
	ErrInvalidInput   = NewAppError("INVALID_INPUT", "Invalid input data")
	ErrUnauthorized   = NewAppError("UNAUTHORIZED", "Unauthorized access")
	ErrForbidden      = NewAppError("FORBIDDEN", "Access forbidden")
	ErrInternalServer = NewAppError("INTERNAL_SERVER", "Internal server error")
	ErrConflict       = NewAppError("CONFLICT", "Resource conflict")
)
