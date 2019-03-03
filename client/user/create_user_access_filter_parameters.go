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

	models "looker-go-sdk/models"
)

// NewCreateUserAccessFilterParams creates a new CreateUserAccessFilterParams object
// with the default values initialized.
func NewCreateUserAccessFilterParams() *CreateUserAccessFilterParams {
	var ()
	return &CreateUserAccessFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserAccessFilterParamsWithTimeout creates a new CreateUserAccessFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserAccessFilterParamsWithTimeout(timeout time.Duration) *CreateUserAccessFilterParams {
	var ()
	return &CreateUserAccessFilterParams{

		timeout: timeout,
	}
}

// NewCreateUserAccessFilterParamsWithContext creates a new CreateUserAccessFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserAccessFilterParamsWithContext(ctx context.Context) *CreateUserAccessFilterParams {
	var ()
	return &CreateUserAccessFilterParams{

		Context: ctx,
	}
}

// NewCreateUserAccessFilterParamsWithHTTPClient creates a new CreateUserAccessFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserAccessFilterParamsWithHTTPClient(client *http.Client) *CreateUserAccessFilterParams {
	var ()
	return &CreateUserAccessFilterParams{
		HTTPClient: client,
	}
}

/*CreateUserAccessFilterParams contains all the parameters to send to the API endpoint
for the create user access filter operation typically these are written to a http.Request
*/
type CreateUserAccessFilterParams struct {

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

// WithTimeout adds the timeout to the create user access filter params
func (o *CreateUserAccessFilterParams) WithTimeout(timeout time.Duration) *CreateUserAccessFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user access filter params
func (o *CreateUserAccessFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user access filter params
func (o *CreateUserAccessFilterParams) WithContext(ctx context.Context) *CreateUserAccessFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user access filter params
func (o *CreateUserAccessFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user access filter params
func (o *CreateUserAccessFilterParams) WithHTTPClient(client *http.Client) *CreateUserAccessFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user access filter params
func (o *CreateUserAccessFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user access filter params
func (o *CreateUserAccessFilterParams) WithBody(body *models.AccessFilter) *CreateUserAccessFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user access filter params
func (o *CreateUserAccessFilterParams) SetBody(body *models.AccessFilter) {
	o.Body = body
}

// WithFields adds the fields to the create user access filter params
func (o *CreateUserAccessFilterParams) WithFields(fields *string) *CreateUserAccessFilterParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create user access filter params
func (o *CreateUserAccessFilterParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the create user access filter params
func (o *CreateUserAccessFilterParams) WithUserID(userID int64) *CreateUserAccessFilterParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create user access filter params
func (o *CreateUserAccessFilterParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserAccessFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
