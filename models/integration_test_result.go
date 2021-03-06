// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IntegrationTestResult integration test result
// swagger:model IntegrationTestResult
type IntegrationTestResult struct {

	// A message representing the results of the test.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Whether or not the test was successful
	// Read Only: true
	Success *bool `json:"success,omitempty"`
}

// Validate validates this integration test result
func (m *IntegrationTestResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationTestResult) UnmarshalBinary(b []byte) error {
	var res IntegrationTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
