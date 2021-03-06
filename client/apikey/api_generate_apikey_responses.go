// Code generated by go-swagger; DO NOT EDIT.

package apikey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jbe-dw/go-powerdns-admin/models"
)

// APIGenerateApikeyReader is a Reader for the APIGenerateApikey structure.
type APIGenerateApikeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIGenerateApikeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAPIGenerateApikeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAPIGenerateApikeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIGenerateApikeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIGenerateApikeyCreated creates a APIGenerateApikeyCreated with default headers values
func NewAPIGenerateApikeyCreated() *APIGenerateApikeyCreated {
	return &APIGenerateApikeyCreated{}
}

/*APIGenerateApikeyCreated handles this case with default header values.

Created
*/
type APIGenerateApikeyCreated struct {
	Payload *models.APIKey
}

func (o *APIGenerateApikeyCreated) Error() string {
	return fmt.Sprintf("[POST /pdnsadmin/apikeys][%d] apiGenerateApikeyCreated  %+v", 201, o.Payload)
}

func (o *APIGenerateApikeyCreated) GetPayload() *models.APIKey {
	return o.Payload
}

func (o *APIGenerateApikeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGenerateApikeyUnprocessableEntity creates a APIGenerateApikeyUnprocessableEntity with default headers values
func NewAPIGenerateApikeyUnprocessableEntity() *APIGenerateApikeyUnprocessableEntity {
	return &APIGenerateApikeyUnprocessableEntity{}
}

/*APIGenerateApikeyUnprocessableEntity handles this case with default header values.

Unprocessable Entry, the ApiKey provided has issues.
*/
type APIGenerateApikeyUnprocessableEntity struct {
	Payload *models.Error
}

func (o *APIGenerateApikeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pdnsadmin/apikeys][%d] apiGenerateApikeyUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *APIGenerateApikeyUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGenerateApikeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGenerateApikeyInternalServerError creates a APIGenerateApikeyInternalServerError with default headers values
func NewAPIGenerateApikeyInternalServerError() *APIGenerateApikeyInternalServerError {
	return &APIGenerateApikeyInternalServerError{}
}

/*APIGenerateApikeyInternalServerError handles this case with default header values.

Internal Server Error. There was a problem creating the key
*/
type APIGenerateApikeyInternalServerError struct {
	Payload *models.Error
}

func (o *APIGenerateApikeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pdnsadmin/apikeys][%d] apiGenerateApikeyInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGenerateApikeyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGenerateApikeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
