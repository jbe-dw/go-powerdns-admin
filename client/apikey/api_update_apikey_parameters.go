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

	"github.com/jbe-dw/go-powerdns-admin/models"
)

// NewAPIUpdateApikeyParams creates a new APIUpdateApikeyParams object
// with the default values initialized.
func NewAPIUpdateApikeyParams() *APIUpdateApikeyParams {
	var ()
	return &APIUpdateApikeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIUpdateApikeyParamsWithTimeout creates a new APIUpdateApikeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIUpdateApikeyParamsWithTimeout(timeout time.Duration) *APIUpdateApikeyParams {
	var ()
	return &APIUpdateApikeyParams{

		timeout: timeout,
	}
}

// NewAPIUpdateApikeyParamsWithContext creates a new APIUpdateApikeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIUpdateApikeyParamsWithContext(ctx context.Context) *APIUpdateApikeyParams {
	var ()
	return &APIUpdateApikeyParams{

		Context: ctx,
	}
}

// NewAPIUpdateApikeyParamsWithHTTPClient creates a new APIUpdateApikeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIUpdateApikeyParamsWithHTTPClient(client *http.Client) *APIUpdateApikeyParams {
	var ()
	return &APIUpdateApikeyParams{
		HTTPClient: client,
	}
}

/*APIUpdateApikeyParams contains all the parameters to send to the API endpoint
for the api update apikey operation typically these are written to a http.Request
*/
type APIUpdateApikeyParams struct {

	/*Apikey
	  ApiKey object with the new values

	*/
	Apikey *models.APIKey
	/*ApikeyID
	  The id of the apikey to retrieve

	*/
	ApikeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api update apikey params
func (o *APIUpdateApikeyParams) WithTimeout(timeout time.Duration) *APIUpdateApikeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api update apikey params
func (o *APIUpdateApikeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api update apikey params
func (o *APIUpdateApikeyParams) WithContext(ctx context.Context) *APIUpdateApikeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api update apikey params
func (o *APIUpdateApikeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api update apikey params
func (o *APIUpdateApikeyParams) WithHTTPClient(client *http.Client) *APIUpdateApikeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api update apikey params
func (o *APIUpdateApikeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the api update apikey params
func (o *APIUpdateApikeyParams) WithApikey(apikey *models.APIKey) *APIUpdateApikeyParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the api update apikey params
func (o *APIUpdateApikeyParams) SetApikey(apikey *models.APIKey) {
	o.Apikey = apikey
}

// WithApikeyID adds the apikeyID to the api update apikey params
func (o *APIUpdateApikeyParams) WithApikeyID(apikeyID int64) *APIUpdateApikeyParams {
	o.SetApikeyID(apikeyID)
	return o
}

// SetApikeyID adds the apikeyId to the api update apikey params
func (o *APIUpdateApikeyParams) SetApikeyID(apikeyID int64) {
	o.ApikeyID = apikeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIUpdateApikeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Apikey != nil {
		if err := r.SetBodyParam(o.Apikey); err != nil {
			return err
		}
	}

	// path param apikey_id
	if err := r.SetPathParam("apikey_id", swag.FormatInt64(o.ApikeyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
