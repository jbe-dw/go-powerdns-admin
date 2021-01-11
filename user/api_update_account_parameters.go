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
	"github.com/go-openapi/swag"

	"models"
)

// NewAPIUpdateAccountParams creates a new APIUpdateAccountParams object
// with the default values initialized.
func NewAPIUpdateAccountParams() *APIUpdateAccountParams {
	var ()
	return &APIUpdateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIUpdateAccountParamsWithTimeout creates a new APIUpdateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIUpdateAccountParamsWithTimeout(timeout time.Duration) *APIUpdateAccountParams {
	var ()
	return &APIUpdateAccountParams{

		timeout: timeout,
	}
}

// NewAPIUpdateAccountParamsWithContext creates a new APIUpdateAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIUpdateAccountParamsWithContext(ctx context.Context) *APIUpdateAccountParams {
	var ()
	return &APIUpdateAccountParams{

		Context: ctx,
	}
}

// NewAPIUpdateAccountParamsWithHTTPClient creates a new APIUpdateAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIUpdateAccountParamsWithHTTPClient(client *http.Client) *APIUpdateAccountParams {
	var ()
	return &APIUpdateAccountParams{
		HTTPClient: client,
	}
}

/*APIUpdateAccountParams contains all the parameters to send to the API endpoint
for the api update account operation typically these are written to a http.Request
*/
type APIUpdateAccountParams struct {

	/*Account*/
	Account *models.APIUpdateAccountParamsBody
	/*AccountID
	  The id of the account to modify or delete

	*/
	AccountID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api update account params
func (o *APIUpdateAccountParams) WithTimeout(timeout time.Duration) *APIUpdateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api update account params
func (o *APIUpdateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api update account params
func (o *APIUpdateAccountParams) WithContext(ctx context.Context) *APIUpdateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api update account params
func (o *APIUpdateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api update account params
func (o *APIUpdateAccountParams) WithHTTPClient(client *http.Client) *APIUpdateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api update account params
func (o *APIUpdateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the api update account params
func (o *APIUpdateAccountParams) WithAccount(account *models.APIUpdateAccountParamsBody) *APIUpdateAccountParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the api update account params
func (o *APIUpdateAccountParams) SetAccount(account *models.APIUpdateAccountParamsBody) {
	o.Account = account
}

// WithAccountID adds the accountID to the api update account params
func (o *APIUpdateAccountParams) WithAccountID(accountID int64) *APIUpdateAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the api update account params
func (o *APIUpdateAccountParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *APIUpdateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {
		if err := r.SetBodyParam(o.Account); err != nil {
			return err
		}
	}

	// path param account_id
	if err := r.SetPathParam("account_id", swag.FormatInt64(o.AccountID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
