// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PowerZoneRanges power zone ranges
//
// swagger:model powerZoneRanges
type PowerZoneRanges struct {

	// zones
	Zones ZoneRanges `json:"zones,omitempty"`
}

// Validate validates this power zone ranges
func (m *PowerZoneRanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerZoneRanges) validateZones(formats strfmt.Registry) error {
	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	if err := m.Zones.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("zones")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("zones")
		}
		return err
	}

	return nil
}

// ContextValidate validate this power zone ranges based on the context it is used
func (m *PowerZoneRanges) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerZoneRanges) contextValidateZones(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Zones.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("zones")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("zones")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerZoneRanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerZoneRanges) UnmarshalBinary(b []byte) error {
	var res PowerZoneRanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
