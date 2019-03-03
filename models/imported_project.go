// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ImportedProject imported project
// swagger:model ImportedProject
type ImportedProject struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Flag signifying if a dependency is remote or local
	// Read Only: true
	IsRemote *bool `json:"is_remote,omitempty"`

	// Dependency name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Ref for a remote dependency
	// Read Only: true
	Ref string `json:"ref,omitempty"`

	// Url for a remote dependency
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this imported project
func (m *ImportedProject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportedProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportedProject) UnmarshalBinary(b []byte) error {
	var res ImportedProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
