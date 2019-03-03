// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Dashboard dashboard
// swagger:model Dashboard
type Dashboard struct {

	// Background color
	BackgroundColor string `json:"background_color,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Content Favorite Id
	// Read Only: true
	ContentFavoriteID int64 `json:"content_favorite_id,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// Time that the Dashboard was created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Elements
	// Read Only: true
	DashboardElements []*DashboardElement `json:"dashboard_elements"`

	// Filters
	// Read Only: true
	DashboardFilters []*DashboardFilter `json:"dashboard_filters"`

	// Layouts
	// Read Only: true
	DashboardLayouts []*DashboardLayout `json:"dashboard_layouts"`

	// Whether or not a dashboard is deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Time that the Dashboard was deleted.
	// Read Only: true
	// Format: date-time
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	// Id of User that deleted the dashboard.
	// Read Only: true
	DeleterID int64 `json:"deleter_id,omitempty"`

	// Description
	Description string `json:"description,omitempty"`

	// Relative path of URI of LookML file to edit the dashboard (LookML dashboard only).
	// Read Only: true
	// Format: uri
	EditURI strfmt.URI `json:"edit_uri,omitempty"`

	// Number of times favorited
	// Read Only: true
	FavoriteCount int64 `json:"favorite_count,omitempty"`

	// Is Hidden
	Hidden bool `json:"hidden,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Time the dashboard was last accessed
	// Read Only: true
	// Format: date-time
	LastAccessedAt strfmt.DateTime `json:"last_accessed_at,omitempty"`

	// Id of User that last updated the dashboard.
	// Read Only: true
	LastUpdaterID int64 `json:"last_updater_id,omitempty"`

	// Time last viewed in the Looker web UI
	// Read Only: true
	// Format: date-time
	LastViewedAt strfmt.DateTime `json:"last_viewed_at,omitempty"`

	// configuration option that governs how dashboard loading will happen.
	LoadConfiguration string `json:"load_configuration,omitempty"`

	// Links this dashboard to a particular LookML dashboard such that calling a **sync** operation on that LookML dashboard will update this dashboard to match.
	// Read Only: true
	LookmlLinkID string `json:"lookml_link_id,omitempty"`

	// Model
	// Read Only: true
	Model *LookModel `json:"model,omitempty"`

	// Timezone in which the Dashboard will run by default.
	QueryTimezone string `json:"query_timezone,omitempty"`

	// Is Read-only
	// Read Only: true
	Readonly *bool `json:"readonly,omitempty"`

	// Refresh Interval
	RefreshInterval string `json:"refresh_interval,omitempty"`

	// Refresh Interval as Integer
	// Read Only: true
	RefreshIntervalToI int64 `json:"refresh_interval_to_i,omitempty"`

	// Show filters bar.  **Security Note:** This property only affects the *cosmetic* appearance of the dashboard, not a user's ability to access data. Hiding the filters bar does **NOT** prevent users from changing filters by other means. For information on how to set up secure data access control policies, see [Control User Access to Data](https://docs.looker.com/admin-options/tutorials/permissions#control_user_access_to_data)
	ShowFiltersBar bool `json:"show_filters_bar,omitempty"`

	// Show title
	ShowTitle bool `json:"show_title,omitempty"`

	// Content Metadata Slug
	Slug string `json:"slug,omitempty"`

	// Space
	// Read Only: true
	Space *SpaceBase `json:"space,omitempty"`

	// Id of Space
	SpaceID string `json:"space_id,omitempty"`

	// Color of text on text tiles
	TextTileTextColor string `json:"text_tile_text_color,omitempty"`

	// Tile background color
	TileBackgroundColor string `json:"tile_background_color,omitempty"`

	// Tile text color
	TileTextColor string `json:"tile_text_color,omitempty"`

	// Dashboard Title
	Title string `json:"title,omitempty"`

	// Title color
	TitleColor string `json:"title_color,omitempty"`

	// Id of User
	// Read Only: true
	UserID int64 `json:"user_id,omitempty"`

	// Number of times viewed in the Looker web UI
	// Read Only: true
	ViewCount int64 `json:"view_count,omitempty"`
}

// Validate validates this dashboard
func (m *Dashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardLayouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditURI(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dashboard) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Dashboard) validateDashboardElements(formats strfmt.Registry) error {

	if swag.IsZero(m.DashboardElements) { // not required
		return nil
	}

	for i := 0; i < len(m.DashboardElements); i++ {
		if swag.IsZero(m.DashboardElements[i]) { // not required
			continue
		}

		if m.DashboardElements[i] != nil {
			if err := m.DashboardElements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboard_elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dashboard) validateDashboardFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.DashboardFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.DashboardFilters); i++ {
		if swag.IsZero(m.DashboardFilters[i]) { // not required
			continue
		}

		if m.DashboardFilters[i] != nil {
			if err := m.DashboardFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboard_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dashboard) validateDashboardLayouts(formats strfmt.Registry) error {

	if swag.IsZero(m.DashboardLayouts) { // not required
		return nil
	}

	for i := 0; i < len(m.DashboardLayouts); i++ {
		if swag.IsZero(m.DashboardLayouts[i]) { // not required
			continue
		}

		if m.DashboardLayouts[i] != nil {
			if err := m.DashboardLayouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboard_layouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dashboard) validateDeletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Dashboard) validateEditURI(formats strfmt.Registry) error {

	if swag.IsZero(m.EditURI) { // not required
		return nil
	}

	if err := validate.FormatOf("edit_uri", "body", "uri", m.EditURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Dashboard) validateLastAccessedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastAccessedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_accessed_at", "body", "date-time", m.LastAccessedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Dashboard) validateLastViewedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastViewedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_viewed_at", "body", "date-time", m.LastViewedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Dashboard) validateModel(formats strfmt.Registry) error {

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

func (m *Dashboard) validateSpace(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Dashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dashboard) UnmarshalBinary(b []byte) error {
	var res Dashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
