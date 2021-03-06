// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LDAPConfigTestIssue l d a p config test issue
// swagger:model LDAPConfigTestIssue
type LDAPConfigTestIssue struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Message describing the issue
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Severity of the issue. Error or Warning
	// Read Only: true
	Severity string `json:"severity,omitempty"`
}

// Validate validates this l d a p config test issue
func (m *LDAPConfigTestIssue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LDAPConfigTestIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPConfigTestIssue) UnmarshalBinary(b []byte) error {
	var res LDAPConfigTestIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
