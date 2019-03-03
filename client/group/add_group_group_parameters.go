// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewAddGroupGroupParams creates a new AddGroupGroupParams object
// with the default values initialized.
func NewAddGroupGroupParams() *AddGroupGroupParams {
	var ()
	return &AddGroupGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGroupGroupParamsWithTimeout creates a new AddGroupGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGroupGroupParamsWithTimeout(timeout time.Duration) *AddGroupGroupParams {
	var ()
	return &AddGroupGroupParams{

		timeout: timeout,
	}
}

// NewAddGroupGroupParamsWithContext creates a new AddGroupGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGroupGroupParamsWithContext(ctx context.Context) *AddGroupGroupParams {
	var ()
	return &AddGroupGroupParams{

		Context: ctx,
	}
}

// NewAddGroupGroupParamsWithHTTPClient creates a new AddGroupGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGroupGroupParamsWithHTTPClient(client *http.Client) *AddGroupGroupParams {
	var ()
	return &AddGroupGroupParams{
		HTTPClient: client,
	}
}

/*AddGroupGroupParams contains all the parameters to send to the API endpoint
for the add group group operation typically these are written to a http.Request
*/
type AddGroupGroupParams struct {

	/*Body
	  Group id to add

	*/
	Body *models.GroupIDForGroupInclusion
	/*GroupID
	  Id of group

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add group group params
func (o *AddGroupGroupParams) WithTimeout(timeout time.Duration) *AddGroupGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add group group params
func (o *AddGroupGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add group group params
func (o *AddGroupGroupParams) WithContext(ctx context.Context) *AddGroupGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add group group params
func (o *AddGroupGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add group group params
func (o *AddGroupGroupParams) WithHTTPClient(client *http.Client) *AddGroupGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add group group params
func (o *AddGroupGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add group group params
func (o *AddGroupGroupParams) WithBody(body *models.GroupIDForGroupInclusion) *AddGroupGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add group group params
func (o *AddGroupGroupParams) SetBody(body *models.GroupIDForGroupInclusion) {
	o.Body = body
}

// WithGroupID adds the groupID to the add group group params
func (o *AddGroupGroupParams) WithGroupID(groupID int64) *AddGroupGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the add group group params
func (o *AddGroupGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *AddGroupGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
