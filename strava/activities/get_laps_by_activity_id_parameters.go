// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetLapsByActivityIDParams creates a new GetLapsByActivityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLapsByActivityIDParams() *GetLapsByActivityIDParams {
	return &GetLapsByActivityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLapsByActivityIDParamsWithTimeout creates a new GetLapsByActivityIDParams object
// with the ability to set a timeout on a request.
func NewGetLapsByActivityIDParamsWithTimeout(timeout time.Duration) *GetLapsByActivityIDParams {
	return &GetLapsByActivityIDParams{
		timeout: timeout,
	}
}

// NewGetLapsByActivityIDParamsWithContext creates a new GetLapsByActivityIDParams object
// with the ability to set a context for a request.
func NewGetLapsByActivityIDParamsWithContext(ctx context.Context) *GetLapsByActivityIDParams {
	return &GetLapsByActivityIDParams{
		Context: ctx,
	}
}

// NewGetLapsByActivityIDParamsWithHTTPClient creates a new GetLapsByActivityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLapsByActivityIDParamsWithHTTPClient(client *http.Client) *GetLapsByActivityIDParams {
	return &GetLapsByActivityIDParams{
		HTTPClient: client,
	}
}

/* GetLapsByActivityIDParams contains all the parameters to send to the API endpoint
   for the get laps by activity Id operation.

   Typically these are written to a http.Request.
*/
type GetLapsByActivityIDParams struct {

	/* ID.

	   The identifier of the activity.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get laps by activity Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLapsByActivityIDParams) WithDefaults() *GetLapsByActivityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get laps by activity Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLapsByActivityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) WithTimeout(timeout time.Duration) *GetLapsByActivityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) WithContext(ctx context.Context) *GetLapsByActivityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) WithHTTPClient(client *http.Client) *GetLapsByActivityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) WithID(id int64) *GetLapsByActivityIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get laps by activity Id params
func (o *GetLapsByActivityIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLapsByActivityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}