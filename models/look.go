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

// Look look
// swagger:model Look
type Look struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Content Favorite Id
	// Read Only: true
	ContentFavoriteID int64 `json:"content_favorite_id,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// Time that the Look was created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Whether or not a look is deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Time that the Look was deleted.
	// Read Only: true
	// Format: date-time
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	// Id of User that deleted the look.
	// Read Only: true
	DeleterID int64 `json:"deleter_id,omitempty"`

	// Description
	Description string `json:"description,omitempty"`

	// Embed Url
	// Read Only: true
	EmbedURL string `json:"embed_url,omitempty"`

	// Excel File Url
	// Read Only: true
	ExcelFileURL string `json:"excel_file_url,omitempty"`

	// Number of times favorited
	// Read Only: true
	FavoriteCount int64 `json:"favorite_count,omitempty"`

	// Google Spreadsheet Formula
	// Read Only: true
	GoogleSpreadsheetFormula string `json:"google_spreadsheet_formula,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Image Embed Url
	// Read Only: true
	ImageEmbedURL string `json:"image_embed_url,omitempty"`

	// auto-run query when Look viewed
	IsRunOnLoad bool `json:"is_run_on_load,omitempty"`

	// Time that the Look was last accessed by any user
	// Read Only: true
	// Format: date-time
	LastAccessedAt strfmt.DateTime `json:"last_accessed_at,omitempty"`

	// Id of User that last updated the look.
	// Read Only: true
	LastUpdaterID int64 `json:"last_updater_id,omitempty"`

	// Time last viewed in the Looker web UI
	// Read Only: true
	// Format: date-time
	LastViewedAt strfmt.DateTime `json:"last_viewed_at,omitempty"`

	// Model
	// Read Only: true
	Model *LookModel `json:"model,omitempty"`

	// Is Public
	// Read Only: true
	Public *bool `json:"public,omitempty"`

	// Public Slug
	// Read Only: true
	PublicSlug string `json:"public_slug,omitempty"`

	// Public Url
	// Read Only: true
	PublicURL string `json:"public_url,omitempty"`

	// Query Id
	QueryID int64 `json:"query_id,omitempty"`

	// Short Url
	// Read Only: true
	ShortURL string `json:"short_url,omitempty"`

	// Space of this Look
	// Read Only: true
	Space *SpaceBase `json:"space,omitempty"`

	// Space Id
	SpaceID string `json:"space_id,omitempty"`

	// Look Title
	Title string `json:"title,omitempty"`

	// Time that the Look was updated.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// User
	// Read Only: true
	User *UserIDOnly `json:"user,omitempty"`

	// User Id
	UserID int64 `json:"user_id,omitempty"`

	// Number of times viewed in the Looker web UI
	// Read Only: true
	ViewCount int64 `json:"view_count,omitempty"`
}

// Validate validates this look
func (m *Look) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAccessedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastViewedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Look) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateDeletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateLastAccessedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastAccessedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_accessed_at", "body", "date-time", m.LastAccessedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateLastViewedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastViewedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_viewed_at", "body", "date-time", m.LastViewedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *Look) validateSpace(formats strfmt.Registry) error {

	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

func (m *Look) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Look) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Look) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Look) UnmarshalBinary(b []byte) error {
	var res Look
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}