// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PrefetchDashboardFilterValue prefetch dashboard filter value
// swagger:model PrefetchDashboardFilterValue
type PrefetchDashboardFilterValue struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Dashboard filter name.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Dashboard filter value
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this prefetch dashboard filter value
func (m *PrefetchDashboardFilterValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrefetchDashboardFilterValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrefetchDashboardFilterValue) UnmarshalBinary(b []byte) error {
	var res PrefetchDashboardFilterValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
