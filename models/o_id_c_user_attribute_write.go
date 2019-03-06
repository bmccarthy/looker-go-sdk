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

// OIDCUserAttributeWrite o ID c user attribute write
// swagger:model OIDCUserAttributeWrite
type OIDCUserAttributeWrite struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Name of User Attribute in OIDC
	Name string `json:"name,omitempty"`

	// Required to be in OIDC assertion for login to be allowed to succeed
	Required bool `json:"required,omitempty"`

	// Link to oidc config
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Looker User Attribute Ids
	UserAttributeIds []int64 `json:"user_attribute_ids"`
}

// Validate validates this o ID c user attribute write
func (m *OIDCUserAttributeWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OIDCUserAttributeWrite) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OIDCUserAttributeWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCUserAttributeWrite) UnmarshalBinary(b []byte) error {
	var res OIDCUserAttributeWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}