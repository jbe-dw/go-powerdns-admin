// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewAPIListAccountsParams creates a new APIListAccountsParams object
// with the default values initialized.
func NewAPIListAccountsParams() *APIListAccountsParams {

	return &APIListAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIListAccountsParamsWithTimeout creates a new APIListAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIListAccountsParamsWithTimeout(timeout time.Duration) *APIListAccountsParams {

	return &APIListAccountsParams{

		timeout: timeout,
	}
}

// NewAPIListAccountsParamsWithContext creates a new APIListAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIListAccountsParamsWithContext(ctx context.Context) *APIListAccountsParams {

	return &APIListAccountsParams{

		Context: ctx,
	}
}

// NewAPIListAccountsParamsWithHTTPClient creates a new APIListAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIListAccountsParamsWithHTTPClient(client *http.Client) *APIListAccountsParams {

	return &APIListAccountsParams{
		HTTPClient: client,
	}
}

/*APIListAccountsParams contains all the parameters to send to the API endpoint
for the api list accounts operation typically these are written to a http.Request
*/
type APIListAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api list accounts params
func (o *APIListAccountsParams) WithTimeout(timeout time.Duration) *APIListAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api list accounts params
func (o *APIListAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api list accounts params
func (o *APIListAccountsParams) WithContext(ctx context.Context) *APIListAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api list accounts params
func (o *APIListAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api list accounts params
func (o *APIListAccountsParams) WithHTTPClient(client *http.Client) *APIListAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api list accounts params
func (o *APIListAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *APIListAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
