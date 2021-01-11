// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteZoneReader is a Reader for the DeleteZone structure.
type DeleteZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteZoneNoContent creates a DeleteZoneNoContent with default headers values
func NewDeleteZoneNoContent() *DeleteZoneNoContent {
	return &DeleteZoneNoContent{}
}

/*DeleteZoneNoContent handles this case with default header values.

Returns 204 No Content on success.
*/
type DeleteZoneNoContent struct {
}

func (o *DeleteZoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}/zones/{zone_id}][%d] deleteZoneNoContent ", 204)
}

func (o *DeleteZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}