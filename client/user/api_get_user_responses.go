// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jbe-dw/go-powerdns-admin/models"
)

// APIGetUserReader is a Reader for the APIGetUser structure.
type APIGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIGetUserOK creates a APIGetUserOK with default headers values
func NewAPIGetUserOK() *APIGetUserOK {
	return &APIGetUserOK{}
}

/*APIGetUserOK handles this case with default header values.

Retrieve a specific User
*/
type APIGetUserOK struct {
	Payload *models.User
}

func (o *APIGetUserOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/users/{username}][%d] apiGetUserOK  %+v", 200, o.Payload)
}

func (o *APIGetUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *APIGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserNotFound creates a APIGetUserNotFound with default headers values
func NewAPIGetUserNotFound() *APIGetUserNotFound {
	return &APIGetUserNotFound{}
}

/*APIGetUserNotFound handles this case with default header values.

Not found. The User with the specified username does not exist
*/
type APIGetUserNotFound struct {
	Payload *models.Error
}

func (o *APIGetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/users/{username}][%d] apiGetUserNotFound  %+v", 404, o.Payload)
}

func (o *APIGetUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserInternalServerError creates a APIGetUserInternalServerError with default headers values
func NewAPIGetUserInternalServerError() *APIGetUserInternalServerError {
	return &APIGetUserInternalServerError{}
}

/*APIGetUserInternalServerError handles this case with default header values.

Internal Server Error, user could not be retrieved. Contains error message
*/
type APIGetUserInternalServerError struct {
	Payload *models.Error
}

func (o *APIGetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/users/{username}][%d] apiGetUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGetUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
