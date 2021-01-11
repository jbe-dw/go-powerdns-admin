// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// APIGetAccountByNameReader is a Reader for the APIGetAccountByName structure.
type APIGetAccountByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIGetAccountByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIGetAccountByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIGetAccountByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIGetAccountByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIGetAccountByNameOK creates a APIGetAccountByNameOK with default headers values
func NewAPIGetAccountByNameOK() *APIGetAccountByNameOK {
	return &APIGetAccountByNameOK{}
}

/*APIGetAccountByNameOK handles this case with default header values.

Retrieve a specific account
*/
type APIGetAccountByNameOK struct {
	Payload *models.Account
}

func (o *APIGetAccountByNameOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/accounts/{account_name}][%d] apiGetAccountByNameOK  %+v", 200, o.Payload)
}

func (o *APIGetAccountByNameOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *APIGetAccountByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetAccountByNameNotFound creates a APIGetAccountByNameNotFound with default headers values
func NewAPIGetAccountByNameNotFound() *APIGetAccountByNameNotFound {
	return &APIGetAccountByNameNotFound{}
}

/*APIGetAccountByNameNotFound handles this case with default header values.

Not found. The Account with the specified name does not exist
*/
type APIGetAccountByNameNotFound struct {
	Payload *models.Error
}

func (o *APIGetAccountByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/accounts/{account_name}][%d] apiGetAccountByNameNotFound  %+v", 404, o.Payload)
}

func (o *APIGetAccountByNameNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetAccountByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetAccountByNameInternalServerError creates a APIGetAccountByNameInternalServerError with default headers values
func NewAPIGetAccountByNameInternalServerError() *APIGetAccountByNameInternalServerError {
	return &APIGetAccountByNameInternalServerError{}
}

/*APIGetAccountByNameInternalServerError handles this case with default header values.

Internal Server Error, account could not be retrieved. Contains error message
*/
type APIGetAccountByNameInternalServerError struct {
	Payload *models.Error
}

func (o *APIGetAccountByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/accounts/{account_name}][%d] apiGetAccountByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGetAccountByNameInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetAccountByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}