// Code generated by go-swagger; DO NOT EDIT.

package clubs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/csmarchbanks/stravaql/strava/model"
)

// GetClubActivitiesByIDReader is a Reader for the GetClubActivitiesByID structure.
type GetClubActivitiesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClubActivitiesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClubActivitiesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClubActivitiesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClubActivitiesByIDOK creates a GetClubActivitiesByIDOK with default headers values
func NewGetClubActivitiesByIDOK() *GetClubActivitiesByIDOK {
	return &GetClubActivitiesByIDOK{}
}

/* GetClubActivitiesByIDOK describes a response with status code 200, with default header values.

A list of activities.
*/
type GetClubActivitiesByIDOK struct {
	Payload []*model.SummaryActivity
}

func (o *GetClubActivitiesByIDOK) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/activities][%d] getClubActivitiesByIdOK  %+v", 200, o.Payload)
}
func (o *GetClubActivitiesByIDOK) GetPayload() []*model.SummaryActivity {
	return o.Payload
}

func (o *GetClubActivitiesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClubActivitiesByIDDefault creates a GetClubActivitiesByIDDefault with default headers values
func NewGetClubActivitiesByIDDefault(code int) *GetClubActivitiesByIDDefault {
	return &GetClubActivitiesByIDDefault{
		_statusCode: code,
	}
}

/* GetClubActivitiesByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetClubActivitiesByIDDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get club activities by Id default response
func (o *GetClubActivitiesByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClubActivitiesByIDDefault) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/activities][%d] getClubActivitiesById default  %+v", o._statusCode, o.Payload)
}
func (o *GetClubActivitiesByIDDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetClubActivitiesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
