// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ForgotPasswordCreatedCode is the HTTP code returned for type ForgotPasswordCreated
const ForgotPasswordCreatedCode int = 201

/*
ForgotPasswordCreated forgot password request created

swagger:response forgotPasswordCreated
*/
type ForgotPasswordCreated struct {
}

// NewForgotPasswordCreated creates ForgotPasswordCreated with default headers values
func NewForgotPasswordCreated() *ForgotPasswordCreated {

	return &ForgotPasswordCreated{}
}

// WriteResponse to the client
func (o *ForgotPasswordCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// ForgotPasswordInternalServerErrorCode is the HTTP code returned for type ForgotPasswordInternalServerError
const ForgotPasswordInternalServerErrorCode int = 500

/*
ForgotPasswordInternalServerError internal server error

swagger:response forgotPasswordInternalServerError
*/
type ForgotPasswordInternalServerError struct {
}

// NewForgotPasswordInternalServerError creates ForgotPasswordInternalServerError with default headers values
func NewForgotPasswordInternalServerError() *ForgotPasswordInternalServerError {

	return &ForgotPasswordInternalServerError{}
}

// WriteResponse to the client
func (o *ForgotPasswordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
