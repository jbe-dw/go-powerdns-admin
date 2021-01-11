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

// APIDeleteUserReader is a Reader for the APIDeleteUser structure.
type APIDeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIDeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIDeleteUserNoContent creates a APIDeleteUserNoContent with default headers values
func NewAPIDeleteUserNoContent() *APIDeleteUserNoContent {
	return &APIDeleteUserNoContent{}
}

/*APIDeleteUserNoContent handles this case with default header values.

OK. User is deleted (empty response body)
*/
type APIDeleteUserNoContent struct {
}

func (o *APIDeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/users/{user_id}][%d] apiDeleteUserNoContent ", 204)
}

func (o *APIDeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIDeleteUserNotFound creates a APIDeleteUserNotFound with default headers values
func NewAPIDeleteUserNotFound() *APIDeleteUserNotFound {
	return &APIDeleteUserNotFound{}
}

/*APIDeleteUserNotFound handles this case with default header values.

Not found. The User with the specified user_id does not exist
*/
type APIDeleteUserNotFound struct {
	Payload *models.Error
}

func (o *APIDeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/users/{user_id}][%d] apiDeleteUserNotFound  %+v", 404, o.Payload)
}

func (o *APIDeleteUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIDeleteUserInternalServerError creates a APIDeleteUserInternalServerError with default headers values
func NewAPIDeleteUserInternalServerError() *APIDeleteUserInternalServerError {
	return &APIDeleteUserInternalServerError{}
}

/*APIDeleteUserInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIDeleteUserInternalServerError struct {
	Payload *models.Error
}

func (o *APIDeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/users/{user_id}][%d] apiDeleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIDeleteUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
