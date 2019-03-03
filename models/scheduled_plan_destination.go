// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ScheduledPlanDestination scheduled plan destination
// swagger:model ScheduledPlanDestination
type ScheduledPlanDestination struct {

	// Address for recipient. For email e.g. 'user@example.com'. For webhooks e.g. 'https://domain/path'. For Amazon S3 e.g. 's3://bucket-name/path/'. For SFTP e.g. 'sftp://host-name/path/'.
	Address string `json:"address,omitempty"`

	// Are values formatted? (containing currency symbols, digit separators, etc.
	ApplyFormatting bool `json:"apply_formatting,omitempty"`

	// Whether visualization options are applied to the results.
	ApplyVis bool `json:"apply_vis,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// The data format to send to the given destination. Supported formats vary by destination, but include: "txt", "csv", "inline_json", "json", "json_detail", "xlsx", "html", "wysiwyg_pdf", "assembled_pdf", "wysiwyg_png"
	Format string `json:"format,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Whether the recipient is a Looker user on the current instance (only applicable for email recipients)
	// Read Only: true
	LookerRecipient *bool `json:"looker_recipient,omitempty"`

	// Optional message to be included in scheduled emails
	Message string `json:"message,omitempty"`

	// JSON object containing parameters for external scheduling. For Amazon S3, this requires keys and values for access_key_id and region. For SFTP, this requires a key and value for username.
	Parameters string `json:"parameters,omitempty"`

	// Id of a scheduled plan you own
	ScheduledPlanID int64 `json:"scheduled_plan_id,omitempty"`

	// (Write-Only) JSON object containing secret parameters for external scheduling. For Amazon S3, this requires a key and value for secret_access_key. For SFTP, this requires a key and value for password.
	SecretParameters string `json:"secret_parameters,omitempty"`

	// Type of the address ('email', 'webhook', 's3', or 'sftp')
	Type string `json:"type,omitempty"`
}

// Validate validates this scheduled plan destination
func (m *ScheduledPlanDestination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledPlanDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledPlanDestination) UnmarshalBinary(b []byte) error {
	var res ScheduledPlanDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
