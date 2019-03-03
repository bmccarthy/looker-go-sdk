// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProjectError project error
// swagger:model ProjectError
type ProjectError struct {

	// A stable token that uniquely identifies this class of error, ignoring parameter values. Error message text may vary due to parameters or localization, but error codes do not. For example, a "File not found" error will have the same error code regardless of the filename in question or the user's display language
	// Read Only: true
	Code string `json:"code,omitempty"`

	// The explore associated with this error
	// Read Only: true
	Explore string `json:"explore,omitempty"`

	// The field associated with this error
	// Read Only: true
	FieldName string `json:"field_name,omitempty"`

	// Name of the file containing this error
	// Read Only: true
	FilePath string `json:"file_path,omitempty"`

	// A link to Looker documentation about this error
	// Read Only: true
	HelpURL string `json:"help_url,omitempty"`

	// Error classification: syntax, deprecation, model_configuration, etc
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// Line number in the file of this error
	// Read Only: true
	LineNumber int64 `json:"line_number,omitempty"`

	// Error message which may contain information such as dashboard or model names that may be considered sensitive in some use cases. Avoid storing or sending this message outside of Looker
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The model associated with this error
	// Read Only: true
	ModelID string `json:"model_id,omitempty"`

	// Error parameters
	// Read Only: true
	Params map[string]string `json:"params,omitempty"`

	// A version of the error message that does not contain potentially sensitive information. Suitable for situations in which messages are stored or sent to consumers outside of Looker, such as external logs. Sanitized messages will display "(?)" where sensitive information would appear in the corresponding non-sanitized message
	// Read Only: true
	SanitizedMessage string `json:"sanitized_message,omitempty"`

	// Severity: fatal, error, warning, info, success
	// Read Only: true
	Severity string `json:"severity,omitempty"`
}

// Validate validates this project error
func (m *ProjectError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectError) UnmarshalBinary(b []byte) error {
	var res ProjectError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
