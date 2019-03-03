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
)

// NewDeleteUserAccessFilterParams creates a new DeleteUserAccessFilterParams object
// with the default values initialized.
func NewDeleteUserAccessFilterParams() *DeleteUserAccessFilterParams {
	var ()
	return &DeleteUserAccessFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserAccessFilterParamsWithTimeout creates a new DeleteUserAccessFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserAccessFilterParamsWithTimeout(timeout time.Duration) *DeleteUserAccessFilterParams {
	var ()
	return &DeleteUserAccessFilterParams{

		timeout: timeout,
	}
}

// NewDeleteUserAccessFilterParamsWithContext creates a new DeleteUserAccessFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserAccessFilterParamsWithContext(ctx context.Context) *DeleteUserAccessFilterParams {
	var ()
	return &DeleteUserAccessFilterParams{

		Context: ctx,
	}
}

// NewDeleteUserAccessFilterParamsWithHTTPClient creates a new DeleteUserAccessFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserAccessFilterParamsWithHTTPClient(client *http.Client) *DeleteUserAccessFilterParams {
	var ()
	return &DeleteUserAccessFilterParams{
		HTTPClient: client,
	}
}

/*DeleteUserAccessFilterParams contains all the parameters to send to the API endpoint
for the delete user access filter operation typically these are written to a http.Request
*/
type DeleteUserAccessFilterParams struct {

	/*AccessFilterID
	  id of Access Filter

	*/
	AccessFilterID int64
	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user access filter params
func (o *DeleteUserAccessFilterParams) WithTimeout(timeout time.Duration) *DeleteUserAccessFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user access filter params
func (o *DeleteUserAccessFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user access filter params
func (o *DeleteUserAccessFilterParams) WithContext(ctx context.Context) *DeleteUserAccessFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user access filter params
func (o *DeleteUserAccessFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user access filter params
func (o *DeleteUserAccessFilterParams) WithHTTPClient(client *http.Client) *DeleteUserAccessFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user access filter params
func (o *DeleteUserAccessFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessFilterID adds the accessFilterID to the delete user access filter params
func (o *DeleteUserAccessFilterParams) WithAccessFilterID(accessFilterID int64) *DeleteUserAccessFilterParams {
	o.SetAccessFilterID(accessFilterID)
	return o
}

// SetAccessFilterID adds the accessFilterId to the delete user access filter params
func (o *DeleteUserAccessFilterParams) SetAccessFilterID(accessFilterID int64) {
	o.AccessFilterID = accessFilterID
}

// WithUserID adds the userID to the delete user access filter params
func (o *DeleteUserAccessFilterParams) WithUserID(userID int64) *DeleteUserAccessFilterParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user access filter params
func (o *DeleteUserAccessFilterParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserAccessFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_filter_id
	if err := r.SetPathParam("access_filter_id", swag.FormatInt64(o.AccessFilterID)); err != nil {
		return err
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
