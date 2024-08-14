package common

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound              = errors.New("the requested resource could not be found")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	resp := ErrorResponse{
		Status:  "error",
		Code:    status,
		Message: message,
	}

	err := WriteJSON(w, status, resp, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func ConflictedResourceResponse(w http.ResponseWriter, r *http.Request, message string) {
	errorResponse(w, r, http.StatusConflict, message)
}
