// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewRoleParams creates a new RoleParams object
// with the default values initialized.
func NewRoleParams() *RoleParams {
	var ()
	return &RoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoleParamsWithTimeout creates a new RoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoleParamsWithTimeout(timeout time.Duration) *RoleParams {
	var ()
	return &RoleParams{

		timeout: timeout,
	}
}

// NewRoleParamsWithContext creates a new RoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoleParamsWithContext(ctx context.Context) *RoleParams {
	var ()
	return &RoleParams{

		Context: ctx,
	}
}

// NewRoleParamsWithHTTPClient creates a new RoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRoleParamsWithHTTPClient(client *http.Client) *RoleParams {
	var ()
	return &RoleParams{
		HTTPClient: client,
	}
}

/*RoleParams contains all the parameters to send to the API endpoint
for the role operation typically these are written to a http.Request
*/
type RoleParams struct {

	/*RoleID
	  id of role

	*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the role params
func (o *RoleParams) WithTimeout(timeout time.Duration) *RoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role params
func (o *RoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role params
func (o *RoleParams) WithContext(ctx context.Context) *RoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role params
func (o *RoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role params
func (o *RoleParams) WithHTTPClient(client *http.Client) *RoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role params
func (o *RoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the role params
func (o *RoleParams) WithRoleID(roleID int64) *RoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the role params
func (o *RoleParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
