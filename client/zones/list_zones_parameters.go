// Code generated by go-swagger; DO NOT EDIT.

package zones

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

// NewListZonesParams creates a new ListZonesParams object
// with the default values initialized.
func NewListZonesParams() *ListZonesParams {
	var ()
	return &ListZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListZonesParamsWithTimeout creates a new ListZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListZonesParamsWithTimeout(timeout time.Duration) *ListZonesParams {
	var ()
	return &ListZonesParams{

		timeout: timeout,
	}
}

// NewListZonesParamsWithContext creates a new ListZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListZonesParamsWithContext(ctx context.Context) *ListZonesParams {
	var ()
	return &ListZonesParams{

		Context: ctx,
	}
}

// NewListZonesParamsWithHTTPClient creates a new ListZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListZonesParamsWithHTTPClient(client *http.Client) *ListZonesParams {
	var ()
	return &ListZonesParams{
		HTTPClient: client,
	}
}

/*ListZonesParams contains all the parameters to send to the API endpoint
for the list zones operation typically these are written to a http.Request
*/
type ListZonesParams struct {

	/*ServerID
	  The id of the server to retrieve

	*/
	ServerID string
	/*Zone
	  When set to the name of a zone, only this zone is returned.
	If no zone with that name exists, the response is an empty array.
	This can e.g. be used to check if a zone exists in the database without having to guess/encode the zone's id or to check if a zone exists.


	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list zones params
func (o *ListZonesParams) WithTimeout(timeout time.Duration) *ListZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list zones params
func (o *ListZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list zones params
func (o *ListZonesParams) WithContext(ctx context.Context) *ListZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list zones params
func (o *ListZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list zones params
func (o *ListZonesParams) WithHTTPClient(client *http.Client) *ListZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list zones params
func (o *ListZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServerID adds the serverID to the list zones params
func (o *ListZonesParams) WithServerID(serverID string) *ListZonesParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the list zones params
func (o *ListZonesParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithZone adds the zone to the list zones params
func (o *ListZonesParams) WithZone(zone *string) *ListZonesParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list zones params
func (o *ListZonesParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *ListZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
