// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jbe-dw/go-powerdns-admin/models"
)

// APIRemoveAccountUserReader is a Reader for the APIRemoveAccountUser structure.
type APIRemoveAccountUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIRemoveAccountUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIRemoveAccountUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIRemoveAccountUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIRemoveAccountUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIRemoveAccountUserNoContent creates a APIRemoveAccountUserNoContent with default headers values
func NewAPIRemoveAccountUserNoContent() *APIRemoveAccountUserNoContent {
	return &APIRemoveAccountUserNoContent{}
}

/*APIRemoveAccountUserNoContent handles this case with default header values.

OK. User is unlinked (empty response body)
*/
type APIRemoveAccountUserNoContent struct {
}

func (o *APIRemoveAccountUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiRemoveAccountUserNoContent ", 204)
}

func (o *APIRemoveAccountUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIRemoveAccountUserNotFound creates a APIRemoveAccountUserNotFound with default headers values
func NewAPIRemoveAccountUserNotFound() *APIRemoveAccountUserNotFound {
	return &APIRemoveAccountUserNotFound{}
}

/*APIRemoveAccountUserNotFound handles this case with default header values.

Not found. The Account or User with the specified id does not exist or user was not linked to account
*/
type APIRemoveAccountUserNotFound struct {
	Payload *models.Error
}

func (o *APIRemoveAccountUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiRemoveAccountUserNotFound  %+v", 404, o.Payload)
}

func (o *APIRemoveAccountUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIRemoveAccountUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIRemoveAccountUserInternalServerError creates a APIRemoveAccountUserInternalServerError with default headers values
func NewAPIRemoveAccountUserInternalServerError() *APIRemoveAccountUserInternalServerError {
	return &APIRemoveAccountUserInternalServerError{}
}

/*APIRemoveAccountUserInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIRemoveAccountUserInternalServerError struct {
	Payload *models.Error
}

func (o *APIRemoveAccountUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/users/{account_id}/{user_id}][%d] apiRemoveAccountUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIRemoveAccountUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIRemoveAccountUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
