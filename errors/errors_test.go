package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerBadRequest(t *testing.T) {
	err := HandlerBadRequest("this is test for bad request")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is test for bad request", err.Message())
	assert.EqualValues(t, "message: this is test for bad request - status: 400 - error: bad_request - causes: []", err.Error())
}
func TestHandlerNotFoundError(t *testing.T) {
	err := HandlerNotFoundError("this is test for not found request")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is test for not found request", err.Message())
	assert.EqualValues(t, "message: this is test for not found request - status: 404 - error: not_found - causes: []", err.Error())
}
func TestHandlerInternalServerError(t *testing.T) {
	err := HandlerInternalServerError("this is test for internal server error", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is test for internal server error", err.Message())
	assert.EqualValues(t, "message: this is test for internal server error - status: 500 - error: internal_server_error - causes: [database error]", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, "database error", err.Causes()[0])
}
