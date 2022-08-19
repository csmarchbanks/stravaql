// Code generated by go-swagger; DO NOT EDIT.

package athletes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/csmarchbanks/stravaql/strava/model"
)

// GetLoggedInAthleteReader is a Reader for the GetLoggedInAthlete structure.
type GetLoggedInAthleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggedInAthleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoggedInAthleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLoggedInAthleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLoggedInAthleteOK creates a GetLoggedInAthleteOK with default headers values
func NewGetLoggedInAthleteOK() *GetLoggedInAthleteOK {
	return &GetLoggedInAthleteOK{}
}

/* GetLoggedInAthleteOK describes a response with status code 200, with default header values.

Profile information for the authenticated athlete.
*/
type GetLoggedInAthleteOK struct {
	Payload *model.DetailedAthlete
}

func (o *GetLoggedInAthleteOK) Error() string {
	return fmt.Sprintf("[GET /athlete][%d] getLoggedInAthleteOK  %+v", 200, o.Payload)
}
func (o *GetLoggedInAthleteOK) GetPayload() *model.DetailedAthlete {
	return o.Payload
}

func (o *GetLoggedInAthleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedAthlete)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInAthleteDefault creates a GetLoggedInAthleteDefault with default headers values
func NewGetLoggedInAthleteDefault(code int) *GetLoggedInAthleteDefault {
	return &GetLoggedInAthleteDefault{
		_statusCode: code,
	}
}

/* GetLoggedInAthleteDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetLoggedInAthleteDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get logged in athlete default response
func (o *GetLoggedInAthleteDefault) Code() int {
	return o._statusCode
}

func (o *GetLoggedInAthleteDefault) Error() string {
	return fmt.Sprintf("[GET /athlete][%d] getLoggedInAthlete default  %+v", o._statusCode, o.Payload)
}
func (o *GetLoggedInAthleteDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetLoggedInAthleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}