// Code generated by go-swagger; DO NOT EDIT.

package apikey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// APIDeleteApikeyReader is a Reader for the APIDeleteApikey structure.
type APIDeleteApikeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIDeleteApikeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIDeleteApikeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIDeleteApikeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIDeleteApikeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIDeleteApikeyNoContent creates a APIDeleteApikeyNoContent with default headers values
func NewAPIDeleteApikeyNoContent() *APIDeleteApikeyNoContent {
	return &APIDeleteApikeyNoContent{}
}

/*APIDeleteApikeyNoContent handles this case with default header values.

OK, key was deleted
*/
type APIDeleteApikeyNoContent struct {
}

func (o *APIDeleteApikeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/apikeys/{apikey_id}][%d] apiDeleteApikeyNoContent ", 204)
}

func (o *APIDeleteApikeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIDeleteApikeyNotFound creates a APIDeleteApikeyNotFound with default headers values
func NewAPIDeleteApikeyNotFound() *APIDeleteApikeyNotFound {
	return &APIDeleteApikeyNotFound{}
}

/*APIDeleteApikeyNotFound handles this case with default header values.

Not found. The ApiKey with the specified apikey_id does not exist
*/
type APIDeleteApikeyNotFound struct {
	Payload *models.Error
}

func (o *APIDeleteApikeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/apikeys/{apikey_id}][%d] apiDeleteApikeyNotFound  %+v", 404, o.Payload)
}

func (o *APIDeleteApikeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteApikeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIDeleteApikeyInternalServerError creates a APIDeleteApikeyInternalServerError with default headers values
func NewAPIDeleteApikeyInternalServerError() *APIDeleteApikeyInternalServerError {
	return &APIDeleteApikeyInternalServerError{}
}

/*APIDeleteApikeyInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIDeleteApikeyInternalServerError struct {
	Payload *models.Error
}

func (o *APIDeleteApikeyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/apikeys/{apikey_id}][%d] apiDeleteApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *APIDeleteApikeyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteApikeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}