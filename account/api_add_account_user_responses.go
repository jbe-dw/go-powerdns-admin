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

// APIAddAccountUserReader is a Reader for the APIAddAccountUser structure.
type APIAddAccountUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIAddAccountUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIAddAccountUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIAddAccountUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIAddAccountUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIAddAccountUserNoContent creates a APIAddAccountUserNoContent with default headers values
func NewAPIAddAccountUserNoContent() *APIAddAccountUserNoContent {
	return &APIAddAccountUserNoContent{}
}

/*APIAddAccountUserNoContent handles this case with default header values.

OK. User is linked (empty response body)
*/
type APIAddAccountUserNoContent struct {
}

func (o *APIAddAccountUserNoContent) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiAddAccountUserNoContent ", 204)
}

func (o *APIAddAccountUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIAddAccountUserNotFound creates a APIAddAccountUserNotFound with default headers values
func NewAPIAddAccountUserNotFound() *APIAddAccountUserNotFound {
	return &APIAddAccountUserNotFound{}
}

/*APIAddAccountUserNotFound handles this case with default header values.

Not found. The Account or User with the specified id does not exist
*/
type APIAddAccountUserNotFound struct {
	Payload *models.Error
}

func (o *APIAddAccountUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiAddAccountUserNotFound  %+v", 404, o.Payload)
}

func (o *APIAddAccountUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIAddAccountUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIAddAccountUserInternalServerError creates a APIAddAccountUserInternalServerError with default headers values
func NewAPIAddAccountUserInternalServerError() *APIAddAccountUserInternalServerError {
	return &APIAddAccountUserInternalServerError{}
}

/*APIAddAccountUserInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIAddAccountUserInternalServerError struct {
	Payload *models.Error
}

func (o *APIAddAccountUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiAddAccountUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIAddAccountUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIAddAccountUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}