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

// APIListUsersReader is a Reader for the APIListUsers structure.
type APIListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAPIListUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIListUsersOK creates a APIListUsersOK with default headers values
func NewAPIListUsersOK() *APIListUsersOK {
	return &APIListUsersOK{}
}

/*APIListUsersOK handles this case with default header values.

List of User objects
*/
type APIListUsersOK struct {
	Payload []*models.User
}

func (o *APIListUsersOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/users][%d] apiListUsersOK  %+v", 200, o.Payload)
}

func (o *APIListUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *APIListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersInternalServerError creates a APIListUsersInternalServerError with default headers values
func NewAPIListUsersInternalServerError() *APIListUsersInternalServerError {
	return &APIListUsersInternalServerError{}
}

/*APIListUsersInternalServerError handles this case with default header values.

Internal Server Error, users could not be retrieved. Contains error message
*/
type APIListUsersInternalServerError struct {
	Payload *models.Error
}

func (o *APIListUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/users][%d] apiListUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *APIListUsersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIListUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
