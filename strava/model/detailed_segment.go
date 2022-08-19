// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetailedSegment detailed segment
//
// swagger:model detailedSegment
type DetailedSegment struct {
	SummarySegment

	// The number of unique athletes who have an effort for this segment
	AthleteCount int64 `json:"athlete_count,omitempty"`

	// The time at which the segment was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The total number of efforts for this segment
	EffortCount int64 `json:"effort_count,omitempty"`

	// Whether this segment is considered hazardous
	Hazardous bool `json:"hazardous,omitempty"`

	// map
	Map *PolylineMap `json:"map,omitempty"`

	// The number of stars for this segment
	StarCount int64 `json:"star_count,omitempty"`

	// The segment's total elevation gain.
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

	// The time at which the segment was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DetailedSegment) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SummarySegment
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SummarySegment = aO0

	// AO1
	var dataAO1 struct {
		AthleteCount int64 `json:"athlete_count,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		EffortCount int64 `json:"effort_count,omitempty"`

		Hazardous bool `json:"hazardous,omitempty"`

		Map *PolylineMap `json:"map,omitempty"`

		StarCount int64 `json:"star_count,omitempty"`

		TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AthleteCount = dataAO1.AthleteCount

	m.CreatedAt = dataAO1.CreatedAt

	m.EffortCount = dataAO1.EffortCount

	m.Hazardous = dataAO1.Hazardous

	m.Map = dataAO1.Map

	m.StarCount = dataAO1.StarCount

	m.TotalElevationGain = dataAO1.TotalElevationGain

	m.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DetailedSegment) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SummarySegment)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AthleteCount int64 `json:"athlete_count,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		EffortCount int64 `json:"effort_count,omitempty"`

		Hazardous bool `json:"hazardous,omitempty"`

		Map *PolylineMap `json:"map,omitempty"`

		StarCount int64 `json:"star_count,omitempty"`

		TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
	}

	dataAO1.AthleteCount = m.AthleteCount

	dataAO1.CreatedAt = m.CreatedAt

	dataAO1.EffortCount = m.EffortCount

	dataAO1.Hazardous = m.Hazardous

	dataAO1.Map = m.Map

	dataAO1.StarCount = m.StarCount

	dataAO1.TotalElevationGain = m.TotalElevationGain

	dataAO1.UpdatedAt = m.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this detailed segment
func (m *DetailedSegment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummarySegment
	if err := m.SummarySegment.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedSegment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DetailedSegment) validateMap(formats strfmt.Registry) error {

	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if m.Map != nil {
		if err := m.Map.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedSegment) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this detailed segment based on the context it is used
func (m *DetailedSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummarySegment
	if err := m.SummarySegment.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedSegment) contextValidateMap(ctx context.Context, formats strfmt.Registry) error {

	if m.Map != nil {
		if err := m.Map.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedSegment) UnmarshalBinary(b []byte) error {
	var res DetailedSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
