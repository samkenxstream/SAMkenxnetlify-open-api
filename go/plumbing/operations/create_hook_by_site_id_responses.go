// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// CreateHookBySiteIDReader is a Reader for the CreateHookBySiteID structure.
type CreateHookBySiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHookBySiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateHookBySiteIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateHookBySiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHookBySiteIDCreated creates a CreateHookBySiteIDCreated with default headers values
func NewCreateHookBySiteIDCreated() *CreateHookBySiteIDCreated {
	return &CreateHookBySiteIDCreated{}
}

/*CreateHookBySiteIDCreated handles this case with default header values.

OK
*/
type CreateHookBySiteIDCreated struct {
	Payload *models.Hook
}

func (o *CreateHookBySiteIDCreated) Error() string {
	return fmt.Sprintf("[POST /hooks][%d] createHookBySiteIdCreated  %+v", 201, o.Payload)
}

func (o *CreateHookBySiteIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHookBySiteIDDefault creates a CreateHookBySiteIDDefault with default headers values
func NewCreateHookBySiteIDDefault(code int) *CreateHookBySiteIDDefault {
	return &CreateHookBySiteIDDefault{
		_statusCode: code,
	}
}

/*CreateHookBySiteIDDefault handles this case with default header values.

error
*/
type CreateHookBySiteIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create hook by site Id default response
func (o *CreateHookBySiteIDDefault) Code() int {
	return o._statusCode
}

func (o *CreateHookBySiteIDDefault) Error() string {
	return fmt.Sprintf("[POST /hooks][%d] createHookBySiteId default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHookBySiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
