// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AccessToken access token
// swagger:model AccessToken
type AccessToken struct {

	// Access Token used for API calls
	// Read Only: true
	AccessToken string `json:"access_token,omitempty"`

	// Number of seconds before the token expires
	// Read Only: true
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// Type of Token
	// Read Only: true
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this access token
func (m *AccessToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessToken) UnmarshalBinary(b []byte) error {
	var res AccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
