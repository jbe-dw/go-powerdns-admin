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

	"models"
)

// NewAPILoginCreateZoneParams creates a new APILoginCreateZoneParams object
// with the default values initialized.
func NewAPILoginCreateZoneParams() *APILoginCreateZoneParams {
	var ()
	return &APILoginCreateZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPILoginCreateZoneParamsWithTimeout creates a new APILoginCreateZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPILoginCreateZoneParamsWithTimeout(timeout time.Duration) *APILoginCreateZoneParams {
	var ()
	return &APILoginCreateZoneParams{

		timeout: timeout,
	}
}

// NewAPILoginCreateZoneParamsWithContext creates a new APILoginCreateZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPILoginCreateZoneParamsWithContext(ctx context.Context) *APILoginCreateZoneParams {
	var ()
	return &APILoginCreateZoneParams{

		Context: ctx,
	}
}

// NewAPILoginCreateZoneParamsWithHTTPClient creates a new APILoginCreateZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPILoginCreateZoneParamsWithHTTPClient(client *http.Client) *APILoginCreateZoneParams {
	var ()
	return &APILoginCreateZoneParams{
		HTTPClient: client,
	}
}

/*APILoginCreateZoneParams contains all the parameters to send to the API endpoint
for the api login create zone operation typically these are written to a http.Request
*/
type APILoginCreateZoneParams struct {

	/*ZoneStruct
	  The zone struct to patch with

	*/
	ZoneStruct *models.Zone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api login create zone params
func (o *APILoginCreateZoneParams) WithTimeout(timeout time.Duration) *APILoginCreateZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api login create zone params
func (o *APILoginCreateZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api login create zone params
func (o *APILoginCreateZoneParams) WithContext(ctx context.Context) *APILoginCreateZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api login create zone params
func (o *APILoginCreateZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api login create zone params
func (o *APILoginCreateZoneParams) WithHTTPClient(client *http.Client) *APILoginCreateZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api login create zone params
func (o *APILoginCreateZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZoneStruct adds the zoneStruct to the api login create zone params
func (o *APILoginCreateZoneParams) WithZoneStruct(zoneStruct *models.Zone) *APILoginCreateZoneParams {
	o.SetZoneStruct(zoneStruct)
	return o
}

// SetZoneStruct adds the zoneStruct to the api login create zone params
func (o *APILoginCreateZoneParams) SetZoneStruct(zoneStruct *models.Zone) {
	o.ZoneStruct = zoneStruct
}

// WriteToRequest writes these params to a swagger request
func (o *APILoginCreateZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ZoneStruct != nil {
		if err := r.SetBodyParam(o.ZoneStruct); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}