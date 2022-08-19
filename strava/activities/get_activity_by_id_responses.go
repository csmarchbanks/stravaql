// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/csmarchbanks/stravaql/strava/model"
)

// GetActivityByIDReader is a Reader for the GetActivityByID structure.
type GetActivityByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivityByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActivityByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetActivityByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetActivityByIDOK creates a GetActivityByIDOK with default headers values
func NewGetActivityByIDOK() *GetActivityByIDOK {
	return &GetActivityByIDOK{}
}

/* GetActivityByIDOK describes a response with status code 200, with default header values.

The activity's detailed representation.
*/
type GetActivityByIDOK struct {
	Payload *model.DetailedActivity
}

func (o *GetActivityByIDOK) Error() string {
	return fmt.Sprintf("[GET /activities/{id}][%d] getActivityByIdOK  %+v", 200, o.Payload)
}
func (o *GetActivityByIDOK) GetPayload() *model.DetailedActivity {
	return o.Payload
}

func (o *GetActivityByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActivityByIDDefault creates a GetActivityByIDDefault with default headers values
func NewGetActivityByIDDefault(code int) *GetActivityByIDDefault {
	return &GetActivityByIDDefault{
		_statusCode: code,
	}
}

/* GetActivityByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetActivityByIDDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get activity by Id default response
func (o *GetActivityByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetActivityByIDDefault) Error() string {
	return fmt.Sprintf("[GET /activities/{id}][%d] getActivityById default  %+v", o._statusCode, o.Payload)
}
func (o *GetActivityByIDDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetActivityByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
