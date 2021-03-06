// Code generated by go-swagger; DO NOT EDIT.

package user

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
)

// NewAPIGetUserParams creates a new APIGetUserParams object
// with the default values initialized.
func NewAPIGetUserParams() *APIGetUserParams {
	var ()
	return &APIGetUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIGetUserParamsWithTimeout creates a new APIGetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIGetUserParamsWithTimeout(timeout time.Duration) *APIGetUserParams {
	var ()
	return &APIGetUserParams{

		timeout: timeout,
	}
}

// NewAPIGetUserParamsWithContext creates a new APIGetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIGetUserParamsWithContext(ctx context.Context) *APIGetUserParams {
	var ()
	return &APIGetUserParams{

		Context: ctx,
	}
}

// NewAPIGetUserParamsWithHTTPClient creates a new APIGetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIGetUserParamsWithHTTPClient(client *http.Client) *APIGetUserParams {
	var ()
	return &APIGetUserParams{
		HTTPClient: client,
	}
}

/*APIGetUserParams contains all the parameters to send to the API endpoint
for the api get user operation typically these are written to a http.Request
*/
type APIGetUserParams struct {

	/*Username
	  The username of the user to retrieve

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api get user params
func (o *APIGetUserParams) WithTimeout(timeout time.Duration) *APIGetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api get user params
func (o *APIGetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api get user params
func (o *APIGetUserParams) WithContext(ctx context.Context) *APIGetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api get user params
func (o *APIGetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api get user params
func (o *APIGetUserParams) WithHTTPClient(client *http.Client) *APIGetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api get user params
func (o *APIGetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the api get user params
func (o *APIGetUserParams) WithUsername(username string) *APIGetUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the api get user params
func (o *APIGetUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *APIGetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
