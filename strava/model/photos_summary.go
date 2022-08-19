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

// PhotosSummary photos summary
//
// swagger:model photosSummary
type PhotosSummary struct {

	// The number of photos
	Count int64 `json:"count,omitempty"`

	// primary
	Primary *PhotosSummaryPrimary `json:"primary,omitempty"`
}

// Validate validates this photos summary
func (m *PhotosSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhotosSummary) validatePrimary(formats strfmt.Registry) error {
	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this photos summary based on the context it is used
func (m *PhotosSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhotosSummary) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {
		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhotosSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhotosSummary) UnmarshalBinary(b []byte) error {
	var res PhotosSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PhotosSummaryPrimary photos summary primary
//
// swagger:model PhotosSummaryPrimary
type PhotosSummaryPrimary struct {

	// id
	ID int64 `json:"id,omitempty"`

	// source
	Source int64 `json:"source,omitempty"`

	// unique id
	UniqueID string `json:"unique_id,omitempty"`

	// urls
	Urls map[string]string `json:"urls,omitempty"`
}

// Validate validates this photos summary primary
func (m *PhotosSummaryPrimary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this photos summary primary based on context it is used
func (m *PhotosSummaryPrimary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhotosSummaryPrimary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhotosSummaryPrimary) UnmarshalBinary(b []byte) error {
	var res PhotosSummaryPrimary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
