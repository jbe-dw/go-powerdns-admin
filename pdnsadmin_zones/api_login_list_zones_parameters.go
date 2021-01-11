// Code generated by go-swagger; DO NOT EDIT.

package pdnsadmin_zones

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

// NewAPILoginListZonesParams creates a new APILoginListZonesParams object
// with the default values initialized.
func NewAPILoginListZonesParams() *APILoginListZonesParams {

	return &APILoginListZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPILoginListZonesParamsWithTimeout creates a new APILoginListZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPILoginListZonesParamsWithTimeout(timeout time.Duration) *APILoginListZonesParams {

	return &APILoginListZonesParams{

		timeout: timeout,
	}
}

// NewAPILoginListZonesParamsWithContext creates a new APILoginListZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPILoginListZonesParamsWithContext(ctx context.Context) *APILoginListZonesParams {

	return &APILoginListZonesParams{

		Context: ctx,
	}
}

// NewAPILoginListZonesParamsWithHTTPClient creates a new APILoginListZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPILoginListZonesParamsWithHTTPClient(client *http.Client) *APILoginListZonesParams {

	return &APILoginListZonesParams{
		HTTPClient: client,
	}
}

/*APILoginListZonesParams contains all the parameters to send to the API endpoint
for the api login list zones operation typically these are written to a http.Request
*/
type APILoginListZonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api login list zones params
func (o *APILoginListZonesParams) WithTimeout(timeout time.Duration) *APILoginListZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api login list zones params
func (o *APILoginListZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api login list zones params
func (o *APILoginListZonesParams) WithContext(ctx context.Context) *APILoginListZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api login list zones params
func (o *APILoginListZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api login list zones params
func (o *APILoginListZonesParams) WithHTTPClient(client *http.Client) *APILoginListZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api login list zones params
func (o *APILoginListZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *APILoginListZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
