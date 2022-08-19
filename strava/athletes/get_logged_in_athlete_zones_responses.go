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

// GetLoggedInAthleteZonesReader is a Reader for the GetLoggedInAthleteZones structure.
type GetLoggedInAthleteZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggedInAthleteZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoggedInAthleteZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLoggedInAthleteZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLoggedInAthleteZonesOK creates a GetLoggedInAthleteZonesOK with default headers values
func NewGetLoggedInAthleteZonesOK() *GetLoggedInAthleteZonesOK {
	return &GetLoggedInAthleteZonesOK{}
}

/* GetLoggedInAthleteZonesOK describes a response with status code 200, with default header values.

Heart rate and power zones.
*/
type GetLoggedInAthleteZonesOK struct {
	Payload *model.Zones
}

func (o *GetLoggedInAthleteZonesOK) Error() string {
	return fmt.Sprintf("[GET /athlete/zones][%d] getLoggedInAthleteZonesOK  %+v", 200, o.Payload)
}
func (o *GetLoggedInAthleteZonesOK) GetPayload() *model.Zones {
	return o.Payload
}

func (o *GetLoggedInAthleteZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Zones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInAthleteZonesDefault creates a GetLoggedInAthleteZonesDefault with default headers values
func NewGetLoggedInAthleteZonesDefault(code int) *GetLoggedInAthleteZonesDefault {
	return &GetLoggedInAthleteZonesDefault{
		_statusCode: code,
	}
}

/* GetLoggedInAthleteZonesDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetLoggedInAthleteZonesDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get logged in athlete zones default response
func (o *GetLoggedInAthleteZonesDefault) Code() int {
	return o._statusCode
}

func (o *GetLoggedInAthleteZonesDefault) Error() string {
	return fmt.Sprintf("[GET /athlete/zones][%d] getLoggedInAthleteZones default  %+v", o._statusCode, o.Payload)
}
func (o *GetLoggedInAthleteZonesDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetLoggedInAthleteZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}