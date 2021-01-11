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

// APIDeleteAccountReader is a Reader for the APIDeleteAccount structure.
type APIDeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIDeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIDeleteAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAPIDeleteAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIDeleteAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIDeleteAccountNoContent creates a APIDeleteAccountNoContent with default headers values
func NewAPIDeleteAccountNoContent() *APIDeleteAccountNoContent {
	return &APIDeleteAccountNoContent{}
}

/*APIDeleteAccountNoContent handles this case with default header values.

OK. Account is deleted (empty response body)
*/
type APIDeleteAccountNoContent struct {
}

func (o *APIDeleteAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/{account_id}][%d] apiDeleteAccountNoContent ", 204)
}

func (o *APIDeleteAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIDeleteAccountNotFound creates a APIDeleteAccountNotFound with default headers values
func NewAPIDeleteAccountNotFound() *APIDeleteAccountNotFound {
	return &APIDeleteAccountNotFound{}
}

/*APIDeleteAccountNotFound handles this case with default header values.

Not found. The Account with the specified account_id does not exist
*/
type APIDeleteAccountNotFound struct {
	Payload *models.Error
}

func (o *APIDeleteAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/{account_id}][%d] apiDeleteAccountNotFound  %+v", 404, o.Payload)
}

func (o *APIDeleteAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIDeleteAccountInternalServerError creates a APIDeleteAccountInternalServerError with default headers values
func NewAPIDeleteAccountInternalServerError() *APIDeleteAccountInternalServerError {
	return &APIDeleteAccountInternalServerError{}
}

/*APIDeleteAccountInternalServerError handles this case with default header values.

Internal Server Error. Contains error message
*/
type APIDeleteAccountInternalServerError struct {
	Payload *models.Error
}

func (o *APIDeleteAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pdnsadmin/accounts/{account_id}][%d] apiDeleteAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *APIDeleteAccountInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *APIDeleteAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
