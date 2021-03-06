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

// APIGetApikeyByIDReader is a Reader for the APIGetApikeyByID structure.
type APIGetApikeyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIGetApikeyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIGetApikeyByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIGetApikeyByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIGetApikeyByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIGetApikeyByIDOK creates a APIGetApikeyByIDOK with default headers values
func NewAPIGetApikeyByIDOK() *APIGetApikeyByIDOK {
	return &APIGetApikeyByIDOK{}
}

/*APIGetApikeyByIDOK handles this case with default header values.

OK.
*/
type APIGetApikeyByIDOK struct {
	Payload *models.APIKey
}

func (o *APIGetApikeyByIDOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/apikeys/{apikey_id}][%d] apiGetApikeyByIdOK  %+v", 200, o.Payload)
}

func (o *APIGetApikeyByIDOK) GetPayload() *models.APIKey {
	return o.Payload
}

func (o *APIGetApikeyByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetApikeyByIDNotFound creates a APIGetApikeyByIDNotFound with default headers values
func NewAPIGetApikeyByIDNotFound() *APIGetApikeyByIDNotFound {
	return &APIGetApikeyByIDNotFound{}
}

/*APIGetApikeyByIDNotFound handles this case with default header values.

Not found. The ApiKey with the specified apikey_id does not exist
*/
type APIGetApikeyByIDNotFound struct {
	Payload *models.Error
}

func (o *APIGetApikeyByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/apikeys/{apikey_id}][%d] apiGetApikeyByIdNotFound  %+v", 404, o.Payload)
}

func (o *APIGetApikeyByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetApikeyByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetApikeyByIDInternalServerError creates a APIGetApikeyByIDInternalServerError with default headers values
func NewAPIGetApikeyByIDInternalServerError() *APIGetApikeyByIDInternalServerError {
	return &APIGetApikeyByIDInternalServerError{}
}

/*APIGetApikeyByIDInternalServerError handles this case with default header values.

Internal Server Error, keys could not be retrieved. Contains error message
*/
type APIGetApikeyByIDInternalServerError struct {
	Payload *models.Error
}

func (o *APIGetApikeyByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/apikeys/{apikey_id}][%d] apiGetApikeyByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGetApikeyByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIGetApikeyByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
