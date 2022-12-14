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

// GetLapsByActivityIDReader is a Reader for the GetLapsByActivityID structure.
type GetLapsByActivityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLapsByActivityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLapsByActivityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLapsByActivityIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLapsByActivityIDOK creates a GetLapsByActivityIDOK with default headers values
func NewGetLapsByActivityIDOK() *GetLapsByActivityIDOK {
	return &GetLapsByActivityIDOK{}
}

/* GetLapsByActivityIDOK describes a response with status code 200, with default header values.

Activity Laps.
*/
type GetLapsByActivityIDOK struct {
	Payload []*model.Lap
}

func (o *GetLapsByActivityIDOK) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/laps][%d] getLapsByActivityIdOK  %+v", 200, o.Payload)
}
func (o *GetLapsByActivityIDOK) GetPayload() []*model.Lap {
	return o.Payload
}

func (o *GetLapsByActivityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLapsByActivityIDDefault creates a GetLapsByActivityIDDefault with default headers values
func NewGetLapsByActivityIDDefault(code int) *GetLapsByActivityIDDefault {
	return &GetLapsByActivityIDDefault{
		_statusCode: code,
	}
}

/* GetLapsByActivityIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetLapsByActivityIDDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get laps by activity Id default response
func (o *GetLapsByActivityIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLapsByActivityIDDefault) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/laps][%d] getLapsByActivityId default  %+v", o._statusCode, o.Payload)
}
func (o *GetLapsByActivityIDDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetLapsByActivityIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
