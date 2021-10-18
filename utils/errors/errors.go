package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"mesage"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewStatusNotImplementedError(message string) RestErr {
	return RestErr{
		Message: message,
		Status:  http.StatusNotImplemented,
		Error:   "not_implemented",
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewUnauthorizeError(message string) *RestErr {
	return &RestErr{
		Message: "unable to retrive user info from given access_token",
		Status:  http.StatusUnauthorized,
		Error:   "unautorized",
	}
}
