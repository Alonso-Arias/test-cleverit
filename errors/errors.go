package errors

import (
	"fmt"
)

// CustomError custom error
type CustomError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	InternalCode string `json:"internalCode"`
}

func (e CustomError) SetMessage(msg string) CustomError {
	e.Message = msg
	return e
}
func (e CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

var (
	BadRequest    = CustomError{Message: "BadRequest", Code: 400, InternalCode: "BADREQUEST"}
	Unauthorized  = CustomError{Message: "Unauthorized", Code: 401, InternalCode: "UNAUTHORIZED"}
	NotFound      = CustomError{Message: "NotFound", Code: 404, InternalCode: "NOT_FOUND"}
	InternalError = CustomError{Message: "Error", Code: 500, InternalCode: "INTERNAL_SERVER_ERROR"}

	InvalidToken = CustomError{Message: "Invalid token", Code: 400, InternalCode: "INVALID_TOKEN"}
	ExpiredToken = CustomError{Message: "Expired token", Code: 400, InternalCode: "EXPIRED_TOKEN"}

	LenPassPolicy   = CustomError{Message: "No contain correct length", Code: 400, InternalCode: "WRONG_PASS_LENGTH"}
	UpperPassPolicy = CustomError{Message: "No contain upper characters", Code: 400, InternalCode: "WRONG_PASS_CONTENT_U"}
	LowerPassPolicy = CustomError{Message: "No contain lower characters", Code: 400, InternalCode: "WRONG_PASS_CONTENT_L"}
	DigitPassPolicy = CustomError{Message: "No contain digit", Code: 400, InternalCode: "WRONG_PASS_CONTENT_D"}

	TasksNotFound     = CustomError{Message: "Tasks not found", Code: 404, InternalCode: "TASKS_NOT_FOUND"}
	TasksAlreadySaved = CustomError{Message: "Tasks already saved", Code: 400, InternalCode: "TASKS_ALREADY_SAVED"}
	TaskStateInvalid  = CustomError{Message: "Task state invalid", Code: 400, InternalCode: "TASK_STATE_INVALID"}
)
