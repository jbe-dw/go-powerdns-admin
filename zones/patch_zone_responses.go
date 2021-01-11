// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchZoneReader is a Reader for the PatchZone structure.
type PatchZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchZoneNoContent creates a PatchZoneNoContent with default headers values
func NewPatchZoneNoContent() *PatchZoneNoContent {
	return &PatchZoneNoContent{}
}

/*PatchZoneNoContent handles this case with default header values.

Returns 204 No Content on success.
*/
type PatchZoneNoContent struct {
}

func (o *PatchZoneNoContent) Error() string {
	return fmt.Sprintf("[PATCH /servers/{server_id}/zones/{zone_id}][%d] patchZoneNoContent ", 204)
}

func (o *PatchZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
