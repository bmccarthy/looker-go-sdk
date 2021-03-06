// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

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

// NewDeleteUserAttributeParams creates a new DeleteUserAttributeParams object
// with the default values initialized.
func NewDeleteUserAttributeParams() *DeleteUserAttributeParams {
	var ()
	return &DeleteUserAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserAttributeParamsWithTimeout creates a new DeleteUserAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserAttributeParamsWithTimeout(timeout time.Duration) *DeleteUserAttributeParams {
	var ()
	return &DeleteUserAttributeParams{

		timeout: timeout,
	}
}

// NewDeleteUserAttributeParamsWithContext creates a new DeleteUserAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserAttributeParamsWithContext(ctx context.Context) *DeleteUserAttributeParams {
	var ()
	return &DeleteUserAttributeParams{

		Context: ctx,
	}
}

// NewDeleteUserAttributeParamsWithHTTPClient creates a new DeleteUserAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserAttributeParamsWithHTTPClient(client *http.Client) *DeleteUserAttributeParams {
	var ()
	return &DeleteUserAttributeParams{
		HTTPClient: client,
	}
}

/*DeleteUserAttributeParams contains all the parameters to send to the API endpoint
for the delete user attribute operation typically these are written to a http.Request
*/
type DeleteUserAttributeParams struct {

	/*UserAttributeID
	  Id of user_attribute

	*/
	UserAttributeID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user attribute params
func (o *DeleteUserAttributeParams) WithTimeout(timeout time.Duration) *DeleteUserAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user attribute params
func (o *DeleteUserAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user attribute params
func (o *DeleteUserAttributeParams) WithContext(ctx context.Context) *DeleteUserAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user attribute params
func (o *DeleteUserAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user attribute params
func (o *DeleteUserAttributeParams) WithHTTPClient(client *http.Client) *DeleteUserAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user attribute params
func (o *DeleteUserAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAttributeID adds the userAttributeID to the delete user attribute params
func (o *DeleteUserAttributeParams) WithUserAttributeID(userAttributeID int64) *DeleteUserAttributeParams {
	o.SetUserAttributeID(userAttributeID)
	return o
}

// SetUserAttributeID adds the userAttributeId to the delete user attribute params
func (o *DeleteUserAttributeParams) SetUserAttributeID(userAttributeID int64) {
	o.UserAttributeID = userAttributeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_attribute_id
	if err := r.SetPathParam("user_attribute_id", swag.FormatInt64(o.UserAttributeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
