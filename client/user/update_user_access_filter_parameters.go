// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// NewUpdateUserAccessFilterParams creates a new UpdateUserAccessFilterParams object
// with the default values initialized.
func NewUpdateUserAccessFilterParams() *UpdateUserAccessFilterParams {
	var ()
	return &UpdateUserAccessFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserAccessFilterParamsWithTimeout creates a new UpdateUserAccessFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserAccessFilterParamsWithTimeout(timeout time.Duration) *UpdateUserAccessFilterParams {
	var ()
	return &UpdateUserAccessFilterParams{

		timeout: timeout,
	}
}

// NewUpdateUserAccessFilterParamsWithContext creates a new UpdateUserAccessFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserAccessFilterParamsWithContext(ctx context.Context) *UpdateUserAccessFilterParams {
	var ()
	return &UpdateUserAccessFilterParams{

		Context: ctx,
	}
}

// NewUpdateUserAccessFilterParamsWithHTTPClient creates a new UpdateUserAccessFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserAccessFilterParamsWithHTTPClient(client *http.Client) *UpdateUserAccessFilterParams {
	var ()
	return &UpdateUserAccessFilterParams{
		HTTPClient: client,
	}
}

/*UpdateUserAccessFilterParams contains all the parameters to send to the API endpoint
for the update user access filter operation typically these are written to a http.Request
*/
type UpdateUserAccessFilterParams struct {

	/*AccessFilterID
	  id of Access Filter

	*/
	AccessFilterID int64
	/*Body
	  Access Filter

	*/
	Body *models.AccessFilter
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithTimeout(timeout time.Duration) *UpdateUserAccessFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithContext(ctx context.Context) *UpdateUserAccessFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithHTTPClient(client *http.Client) *UpdateUserAccessFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessFilterID adds the accessFilterID to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithAccessFilterID(accessFilterID int64) *UpdateUserAccessFilterParams {
	o.SetAccessFilterID(accessFilterID)
	return o
}

// SetAccessFilterID adds the accessFilterId to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetAccessFilterID(accessFilterID int64) {
	o.AccessFilterID = accessFilterID
}

// WithBody adds the body to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithBody(body *models.AccessFilter) *UpdateUserAccessFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetBody(body *models.AccessFilter) {
	o.Body = body
}

// WithFields adds the fields to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithFields(fields *string) *UpdateUserAccessFilterParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the update user access filter params
func (o *UpdateUserAccessFilterParams) WithUserID(userID int64) *UpdateUserAccessFilterParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user access filter params
func (o *UpdateUserAccessFilterParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserAccessFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_filter_id
	if err := r.SetPathParam("access_filter_id", swag.FormatInt64(o.AccessFilterID)); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
