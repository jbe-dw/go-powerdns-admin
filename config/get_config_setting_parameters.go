// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewGetConfigSettingParams creates a new GetConfigSettingParams object
// with the default values initialized.
func NewGetConfigSettingParams() *GetConfigSettingParams {
	var ()
	return &GetConfigSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigSettingParamsWithTimeout creates a new GetConfigSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigSettingParamsWithTimeout(timeout time.Duration) *GetConfigSettingParams {
	var ()
	return &GetConfigSettingParams{

		timeout: timeout,
	}
}

// NewGetConfigSettingParamsWithContext creates a new GetConfigSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigSettingParamsWithContext(ctx context.Context) *GetConfigSettingParams {
	var ()
	return &GetConfigSettingParams{

		Context: ctx,
	}
}

// NewGetConfigSettingParamsWithHTTPClient creates a new GetConfigSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigSettingParamsWithHTTPClient(client *http.Client) *GetConfigSettingParams {
	var ()
	return &GetConfigSettingParams{
		HTTPClient: client,
	}
}

/*GetConfigSettingParams contains all the parameters to send to the API endpoint
for the get config setting operation typically these are written to a http.Request
*/
type GetConfigSettingParams struct {

	/*ConfigSettingName
	  The name of the setting to retrieve

	*/
	ConfigSettingName string
	/*ServerID
	  The id of the server to retrieve

	*/
	ServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get config setting params
func (o *GetConfigSettingParams) WithTimeout(timeout time.Duration) *GetConfigSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config setting params
func (o *GetConfigSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config setting params
func (o *GetConfigSettingParams) WithContext(ctx context.Context) *GetConfigSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config setting params
func (o *GetConfigSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get config setting params
func (o *GetConfigSettingParams) WithHTTPClient(client *http.Client) *GetConfigSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config setting params
func (o *GetConfigSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigSettingName adds the configSettingName to the get config setting params
func (o *GetConfigSettingParams) WithConfigSettingName(configSettingName string) *GetConfigSettingParams {
	o.SetConfigSettingName(configSettingName)
	return o
}

// SetConfigSettingName adds the configSettingName to the get config setting params
func (o *GetConfigSettingParams) SetConfigSettingName(configSettingName string) {
	o.ConfigSettingName = configSettingName
}

// WithServerID adds the serverID to the get config setting params
func (o *GetConfigSettingParams) WithServerID(serverID string) *GetConfigSettingParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the get config setting params
func (o *GetConfigSettingParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config_setting_name
	if err := r.SetPathParam("config_setting_name", o.ConfigSettingName); err != nil {
		return err
	}

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
