// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// ListZoneReader is a Reader for the ListZone structure.
type ListZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListZoneOK creates a ListZoneOK with default headers values
func NewListZoneOK() *ListZoneOK {
	return &ListZoneOK{}
}

/*ListZoneOK handles this case with default header values.

A Zone
*/
type ListZoneOK struct {
	Payload *models.Zone
}

func (o *ListZoneOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}/zones/{zone_id}][%d] listZoneOK  %+v", 200, o.Payload)
}

func (o *ListZoneOK) GetPayload() *models.Zone {
	return o.Payload
}

func (o *ListZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
