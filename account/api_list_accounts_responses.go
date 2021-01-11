// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// APIListAccountsReader is a Reader for the APIListAccounts structure.
type APIListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAPIListAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIListAccountsOK creates a APIListAccountsOK with default headers values
func NewAPIListAccountsOK() *APIListAccountsOK {
	return &APIListAccountsOK{}
}

/*APIListAccountsOK handles this case with default header values.

List of Account objects
*/
type APIListAccountsOK struct {
	Payload []*models.Account
}

func (o *APIListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/accounts][%d] apiListAccountsOK  %+v", 200, o.Payload)
}

func (o *APIListAccountsOK) GetPayload() []*models.Account {
	return o.Payload
}

func (o *APIListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListAccountsInternalServerError creates a APIListAccountsInternalServerError with default headers values
func NewAPIListAccountsInternalServerError() *APIListAccountsInternalServerError {
	return &APIListAccountsInternalServerError{}
}

/*APIListAccountsInternalServerError handles this case with default header values.

Internal Server Error, accounts could not be retrieved. Contains error message
*/
type APIListAccountsInternalServerError struct {
	Payload *models.Error
}

func (o *APIListAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/accounts][%d] apiListAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *APIListAccountsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIListAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
