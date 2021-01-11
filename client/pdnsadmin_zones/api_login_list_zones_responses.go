// Code generated by go-swagger; DO NOT EDIT.

package pdnsadmin_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// APILoginListZonesReader is a Reader for the APILoginListZones structure.
type APILoginListZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APILoginListZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPILoginListZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPILoginListZonesOK creates a APILoginListZonesOK with default headers values
func NewAPILoginListZonesOK() *APILoginListZonesOK {
	return &APILoginListZonesOK{}
}

/*APILoginListZonesOK handles this case with default header values.

An array of Zones
*/
type APILoginListZonesOK struct {
	Payload []models.PDNSAdminZones
}

func (o *APILoginListZonesOK) Error() string {
	return fmt.Sprintf("[GET /pdnsadmin/zones][%d] apiLoginListZonesOK  %+v", 200, o.Payload)
}

func (o *APILoginListZonesOK) GetPayload() []models.PDNSAdminZones {
	return o.Payload
}

func (o *APILoginListZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}