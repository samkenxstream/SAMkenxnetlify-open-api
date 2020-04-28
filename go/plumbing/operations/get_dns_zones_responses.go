// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// GetDNSZonesReader is a Reader for the GetDNSZones structure.
type GetDNSZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDNSZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDNSZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDNSZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDNSZonesOK creates a GetDNSZonesOK with default headers values
func NewGetDNSZonesOK() *GetDNSZonesOK {
	return &GetDNSZonesOK{}
}

/*GetDNSZonesOK handles this case with default header values.

get all DNS zones your user account has access to
*/
type GetDNSZonesOK struct {
	Payload models.DNSZones
}

func (o *GetDNSZonesOK) Error() string {
	return fmt.Sprintf("[GET /dns_zones][%d] getDnsZonesOK  %+v", 200, o.Payload)
}

func (o *GetDNSZonesOK) GetPayload() models.DNSZones {
	return o.Payload
}

func (o *GetDNSZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDNSZonesDefault creates a GetDNSZonesDefault with default headers values
func NewGetDNSZonesDefault(code int) *GetDNSZonesDefault {
	return &GetDNSZonesDefault{
		_statusCode: code,
	}
}

/*GetDNSZonesDefault handles this case with default header values.

error
*/
type GetDNSZonesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get Dns zones default response
func (o *GetDNSZonesDefault) Code() int {
	return o._statusCode
}

func (o *GetDNSZonesDefault) Error() string {
	return fmt.Sprintf("[GET /dns_zones][%d] getDnsZones default  %+v", o._statusCode, o.Payload)
}

func (o *GetDNSZonesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDNSZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
