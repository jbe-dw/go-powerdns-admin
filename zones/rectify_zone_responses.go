// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RectifyZoneReader is a Reader for the RectifyZone structure.
type RectifyZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RectifyZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRectifyZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRectifyZoneOK creates a RectifyZoneOK with default headers values
func NewRectifyZoneOK() *RectifyZoneOK {
	return &RectifyZoneOK{}
}

/*RectifyZoneOK handles this case with default header values.

OK
*/
type RectifyZoneOK struct {
	Payload string
}

func (o *RectifyZoneOK) Error() string {
	return fmt.Sprintf("[PUT /servers/{server_id}/zones/{zone_id}/rectify][%d] rectifyZoneOK  %+v", 200, o.Payload)
}

func (o *RectifyZoneOK) GetPayload() string {
	return o.Payload
}

func (o *RectifyZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}