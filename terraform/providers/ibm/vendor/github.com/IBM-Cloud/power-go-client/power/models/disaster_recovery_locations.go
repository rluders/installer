// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DisasterRecoveryLocations disaster recovery locations
//
// swagger:model DisasterRecoveryLocations
type DisasterRecoveryLocations struct {

	// The list of Disaster Recovery Locations
	// Required: true
	DisasterRecoryLocations []*DisasterRecoveryLocation `json:"disasterRecoryLocations"`
}

// Validate validates this disaster recovery locations
func (m *DisasterRecoveryLocations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisasterRecoryLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisasterRecoveryLocations) validateDisasterRecoryLocations(formats strfmt.Registry) error {

	if err := validate.Required("disasterRecoryLocations", "body", m.DisasterRecoryLocations); err != nil {
		return err
	}

	for i := 0; i < len(m.DisasterRecoryLocations); i++ {
		if swag.IsZero(m.DisasterRecoryLocations[i]) { // not required
			continue
		}

		if m.DisasterRecoryLocations[i] != nil {
			if err := m.DisasterRecoryLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disasterRecoryLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disasterRecoryLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this disaster recovery locations based on the context it is used
func (m *DisasterRecoveryLocations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisasterRecoryLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisasterRecoveryLocations) contextValidateDisasterRecoryLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisasterRecoryLocations); i++ {

		if m.DisasterRecoryLocations[i] != nil {
			if err := m.DisasterRecoryLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disasterRecoryLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disasterRecoryLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DisasterRecoveryLocations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisasterRecoveryLocations) UnmarshalBinary(b []byte) error {
	var res DisasterRecoveryLocations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
