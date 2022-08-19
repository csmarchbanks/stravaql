// Code generated by go-swagger; DO NOT EDIT.

package segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/csmarchbanks/stravaql/strava/model"
)

// GetLoggedInAthleteStarredSegmentsReader is a Reader for the GetLoggedInAthleteStarredSegments structure.
type GetLoggedInAthleteStarredSegmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggedInAthleteStarredSegmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoggedInAthleteStarredSegmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLoggedInAthleteStarredSegmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLoggedInAthleteStarredSegmentsOK creates a GetLoggedInAthleteStarredSegmentsOK with default headers values
func NewGetLoggedInAthleteStarredSegmentsOK() *GetLoggedInAthleteStarredSegmentsOK {
	return &GetLoggedInAthleteStarredSegmentsOK{}
}

/* GetLoggedInAthleteStarredSegmentsOK describes a response with status code 200, with default header values.

List of the authenticated athlete's starred segments.
*/
type GetLoggedInAthleteStarredSegmentsOK struct {
	Payload []*model.SummarySegment
}

func (o *GetLoggedInAthleteStarredSegmentsOK) Error() string {
	return fmt.Sprintf("[GET /segments/starred][%d] getLoggedInAthleteStarredSegmentsOK  %+v", 200, o.Payload)
}
func (o *GetLoggedInAthleteStarredSegmentsOK) GetPayload() []*model.SummarySegment {
	return o.Payload
}

func (o *GetLoggedInAthleteStarredSegmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInAthleteStarredSegmentsDefault creates a GetLoggedInAthleteStarredSegmentsDefault with default headers values
func NewGetLoggedInAthleteStarredSegmentsDefault(code int) *GetLoggedInAthleteStarredSegmentsDefault {
	return &GetLoggedInAthleteStarredSegmentsDefault{
		_statusCode: code,
	}
}

/* GetLoggedInAthleteStarredSegmentsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetLoggedInAthleteStarredSegmentsDefault struct {
	_statusCode int

	Payload *model.Fault
}

// Code gets the status code for the get logged in athlete starred segments default response
func (o *GetLoggedInAthleteStarredSegmentsDefault) Code() int {
	return o._statusCode
}

func (o *GetLoggedInAthleteStarredSegmentsDefault) Error() string {
	return fmt.Sprintf("[GET /segments/starred][%d] getLoggedInAthleteStarredSegments default  %+v", o._statusCode, o.Payload)
}
func (o *GetLoggedInAthleteStarredSegmentsDefault) GetPayload() *model.Fault {
	return o.Payload
}

func (o *GetLoggedInAthleteStarredSegmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}