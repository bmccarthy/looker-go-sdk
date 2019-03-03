// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LegacyFeature legacy feature
// swagger:model LegacyFeature
type LegacyFeature struct {

	// Approximate date that this feature will be automatically disabled.
	// Read Only: true
	// Format: date-time
	ApproximateDisableDate strfmt.DateTime `json:"approximate_disable_date,omitempty"`

	// Approximate date that this feature will be removed.
	// Read Only: true
	// Format: date-time
	ApproximateEndOfLifeDate strfmt.DateTime `json:"approximate_end_of_life_date,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Looker version where this feature will be automatically disabled
	// Read Only: true
	DisableOnUpgradeToVersion string `json:"disable_on_upgrade_to_version,omitempty"`

	// Looker version where this feature became a legacy feature
	// Read Only: true
	DisallowedAsOfVersion string `json:"disallowed_as_of_version,omitempty"`

	// URL for documentation about this feature
	// Read Only: true
	DocumentationURL string `json:"documentation_url,omitempty"`

	// Whether this feature is currently enabled
	// Read Only: true
	Enabled *bool `json:"enabled,omitempty"`

	// Whether this feature has been enabled by a user
	EnabledLocally bool `json:"enabled_locally,omitempty"`

	// Future Looker version where this feature will be removed
	// Read Only: true
	EndOfLifeVersion string `json:"end_of_life_version,omitempty"`

	// Whether this legacy feature may have been automatically disabled when upgrading to the current version.
	// Read Only: true
	HasDisabledOnUpgrade *bool `json:"has_disabled_on_upgrade,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Name
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this legacy feature
func (m *LegacyFeature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApproximateDisableDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApproximateEndOfLifeDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyFeature) validateApproximateDisableDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ApproximateDisableDate) { // not required
		return nil
	}

	if err := validate.FormatOf("approximate_disable_date", "body", "date-time", m.ApproximateDisableDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyFeature) validateApproximateEndOfLifeDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ApproximateEndOfLifeDate) { // not required
		return nil
	}

	if err := validate.FormatOf("approximate_end_of_life_date", "body", "date-time", m.ApproximateEndOfLifeDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyFeature) UnmarshalBinary(b []byte) error {
	var res LegacyFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
