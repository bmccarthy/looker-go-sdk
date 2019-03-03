// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LookBasic look basic
// swagger:model LookBasic
type LookBasic struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Look Title
	// Read Only: true
	Title string `json:"title,omitempty"`
}

// Validate validates this look basic
func (m *LookBasic) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LookBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookBasic) UnmarshalBinary(b []byte) error {
	var res LookBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
