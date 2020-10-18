package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerBadRequest(t *testing.T)  {
	err := HandlerBadRequest("this is test for bad request")
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusBadRequest,err.Status)
	assert.EqualValues(t,"this is test for bad request",err.Message)
	assert.EqualValues(t,"bad_request",err.Error)
}
func TestHandlerNotFoundError(t *testing.T)  {
	err := HandlerNotFoundError("this is test for not found request")
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusNotFound,err.Status)
	assert.EqualValues(t,"this is test for not found request",err.Message)
	assert.EqualValues(t,"not_found",err.Error)
}
func TestHandlerInternalServerError(t *testing.T)  {
	err := HandlerInternalServerError("this is test for internal server error",errors.New("database error"))
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.Status)
	assert.EqualValues(t,"this is test for internal server error",err.Message)
	assert.EqualValues(t,"internal_server_error",err.Error)

	assert.NotNil(t,err.Causes)
	assert.EqualValues(t,"database error", err.Causes[0])
}