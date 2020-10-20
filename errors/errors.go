package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ResError interface
type ResError interface {
	Error() string
	Message() string
	Status() int
	Causes() []interface{}
}

// resError struct
type resError struct {
	message string
	status  int
	error   string
	causes  []interface{}
}

func (re resError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %v - causes: %v", re.message, re.status, re.error, re.causes)
}
func (re resError) Message() string {
	return re.message
}
func (re resError) Status() int {
	return re.status
}
func (re resError) Causes() []interface{} {
	return re.causes
}

// HandlerBadRequest func for ResError struct
func HandlerBadRequest(message string) ResError {
	return resError{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

// HandlerNotFoundError func for ResError struct
func HandlerNotFoundError(message string) ResError {
	return resError{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

// HandlerUnauthorizedError func
func HandlerUnauthorizedError(message string) ResError {
	return resError{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

// HandlerInternalServerError func for ResError struct
func HandlerInternalServerError(message string, err error) ResError {
	result := resError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}

// HandleRestError func
func HandleRestError(message string, status int, er string, causes []interface{}) ResError {
	return resError{
		message: message,
		status:  http.StatusUnauthorized,
		error:   er,
		causes:  causes,
	}
}

// HandleRestErrorFromBytes func
func HandleRestErrorFromBytes(message []byte) (ResError, error) {
	var apiErr resError
	if err := json.Unmarshal(message, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}
