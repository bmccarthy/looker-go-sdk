// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IntegrationRequiredField integration required field
// swagger:model IntegrationRequiredField
type IntegrationRequiredField struct {

	// If present, supercedes 'tag' and matches a field that has all of the provided tags.
	// Read Only: true
	AllTags []string `json:"all_tags"`

	// If present, supercedes 'tag' and matches a field that has any of the provided tags.
	// Read Only: true
	AnyTag []string `json:"any_tag"`

	// Matches a field that has this tag.
	// Read Only: true
	Tag string `json:"tag,omitempty"`
}

// Validate validates this integration required field
func (m *IntegrationRequiredField) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationRequiredField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationRequiredField) UnmarshalBinary(b []byte) error {
	var res IntegrationRequiredField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
