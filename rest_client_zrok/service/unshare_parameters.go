// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// NewUnshareParams creates a new UnshareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnshareParams() *UnshareParams {
	return &UnshareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnshareParamsWithTimeout creates a new UnshareParams object
// with the ability to set a timeout on a request.
func NewUnshareParamsWithTimeout(timeout time.Duration) *UnshareParams {
	return &UnshareParams{
		timeout: timeout,
	}
}

// NewUnshareParamsWithContext creates a new UnshareParams object
// with the ability to set a context for a request.
func NewUnshareParamsWithContext(ctx context.Context) *UnshareParams {
	return &UnshareParams{
		Context: ctx,
	}
}

// NewUnshareParamsWithHTTPClient creates a new UnshareParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnshareParamsWithHTTPClient(client *http.Client) *UnshareParams {
	return &UnshareParams{
		HTTPClient: client,
	}
}

/*
UnshareParams contains all the parameters to send to the API endpoint

	for the unshare operation.

	Typically these are written to a http.Request.
*/
type UnshareParams struct {

	// Body.
	Body *rest_model_zrok.UnshareRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unshare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnshareParams) WithDefaults() *UnshareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unshare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnshareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unshare params
func (o *UnshareParams) WithTimeout(timeout time.Duration) *UnshareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unshare params
func (o *UnshareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unshare params
func (o *UnshareParams) WithContext(ctx context.Context) *UnshareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unshare params
func (o *UnshareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unshare params
func (o *UnshareParams) WithHTTPClient(client *http.Client) *UnshareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unshare params
func (o *UnshareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unshare params
func (o *UnshareParams) WithBody(body *rest_model_zrok.UnshareRequest) *UnshareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unshare params
func (o *UnshareParams) SetBody(body *rest_model_zrok.UnshareRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UnshareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
