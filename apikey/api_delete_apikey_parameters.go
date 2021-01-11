// Code generated by go-swagger; DO NOT EDIT.

package apikey

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
	"github.com/go-openapi/swag"
)

// NewAPIDeleteApikeyParams creates a new APIDeleteApikeyParams object
// with the default values initialized.
func NewAPIDeleteApikeyParams() *APIDeleteApikeyParams {
	var ()
	return &APIDeleteApikeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDeleteApikeyParamsWithTimeout creates a new APIDeleteApikeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDeleteApikeyParamsWithTimeout(timeout time.Duration) *APIDeleteApikeyParams {
	var ()
	return &APIDeleteApikeyParams{

		timeout: timeout,
	}
}

// NewAPIDeleteApikeyParamsWithContext creates a new APIDeleteApikeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDeleteApikeyParamsWithContext(ctx context.Context) *APIDeleteApikeyParams {
	var ()
	return &APIDeleteApikeyParams{

		Context: ctx,
	}
}

// NewAPIDeleteApikeyParamsWithHTTPClient creates a new APIDeleteApikeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDeleteApikeyParamsWithHTTPClient(client *http.Client) *APIDeleteApikeyParams {
	var ()
	return &APIDeleteApikeyParams{
		HTTPClient: client,
	}
}

/*APIDeleteApikeyParams contains all the parameters to send to the API endpoint
for the api delete apikey operation typically these are written to a http.Request
*/
type APIDeleteApikeyParams struct {

	/*ApikeyID
	  The id of the apikey to retrieve

	*/
	ApikeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api delete apikey params
func (o *APIDeleteApikeyParams) WithTimeout(timeout time.Duration) *APIDeleteApikeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api delete apikey params
func (o *APIDeleteApikeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api delete apikey params
func (o *APIDeleteApikeyParams) WithContext(ctx context.Context) *APIDeleteApikeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api delete apikey params
func (o *APIDeleteApikeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api delete apikey params
func (o *APIDeleteApikeyParams) WithHTTPClient(client *http.Client) *APIDeleteApikeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api delete apikey params
func (o *APIDeleteApikeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikeyID adds the apikeyID to the api delete apikey params
func (o *APIDeleteApikeyParams) WithApikeyID(apikeyID int64) *APIDeleteApikeyParams {
	o.SetApikeyID(apikeyID)
	return o
}

// SetApikeyID adds the apikeyId to the api delete apikey params
func (o *APIDeleteApikeyParams) SetApikeyID(apikeyID int64) {
	o.ApikeyID = apikeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIDeleteApikeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apikey_id
	if err := r.SetPathParam("apikey_id", swag.FormatInt64(o.ApikeyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}