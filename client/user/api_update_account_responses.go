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

// APIUpdateAccountReader is a Reader for the APIUpdateAccount structure.
type APIUpdateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIUpdateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIUpdateAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIUpdateAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIUpdateAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIUpdateAccountNoContent creates a APIUpdateAccountNoContent with default headers values
func NewAPIUpdateAccountNoContent() *APIUpdateAccountNoContent {
	return &APIUpdateAccountNoContent{}
}

/*APIUpdateAccountNoContent handles this case with default header values.

OK. Account is modified (empty response body)
*/
type APIUpdateAccountNoContent struct {
}

func (o *APIUpdateAccountNoContent) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/{account_id}][%d] apiUpdateAccountNoContent ", 204)
}

func (o *APIUpdateAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIUpdateAccountNotFound creates a APIUpdateAccountNotFound with default headers values
func NewAPIUpdateAccountNotFound() *APIUpdateAccountNotFound {
	return &APIUpdateAccountNotFound{}
}

/*APIUpdateAccountNotFound handles this case with default header values.

Not found. The Account with the specified account_id does not exist
*/
type APIUpdateAccountNotFound struct {
	Payload *models.Error
}

func (o *APIUpdateAccountNotFound) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/{account_id}][%d] apiUpdateAccountNotFound  %+v", 404, o.Payload)
}

func (o *APIUpdateAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIUpdateAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIUpdateAccountInternalServerError creates a APIUpdateAccountInternalServerError with default headers values
func NewAPIUpdateAccountInternalServerError() *APIUpdateAccountInternalServerError {
	return &APIUpdateAccountInternalServerError{}
}

/*APIUpdateAccountInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIUpdateAccountInternalServerError struct {
	Payload *models.Error
}

func (o *APIUpdateAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/{account_id}][%d] apiUpdateAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *APIUpdateAccountInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIUpdateAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}