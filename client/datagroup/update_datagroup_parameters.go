// Code generated by go-swagger; DO NOT EDIT.

package datagroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// NewUpdateDatagroupParams creates a new UpdateDatagroupParams object
// with the default values initialized.
func NewUpdateDatagroupParams() *UpdateDatagroupParams {
	var ()
	return &UpdateDatagroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDatagroupParamsWithTimeout creates a new UpdateDatagroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDatagroupParamsWithTimeout(timeout time.Duration) *UpdateDatagroupParams {
	var ()
	return &UpdateDatagroupParams{

		timeout: timeout,
	}
}

// NewUpdateDatagroupParamsWithContext creates a new UpdateDatagroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDatagroupParamsWithContext(ctx context.Context) *UpdateDatagroupParams {
	var ()
	return &UpdateDatagroupParams{

		Context: ctx,
	}
}

// NewUpdateDatagroupParamsWithHTTPClient creates a new UpdateDatagroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDatagroupParamsWithHTTPClient(client *http.Client) *UpdateDatagroupParams {
	var ()
	return &UpdateDatagroupParams{
		HTTPClient: client,
	}
}

/*UpdateDatagroupParams contains all the parameters to send to the API endpoint
for the update datagroup operation typically these are written to a http.Request
*/
type UpdateDatagroupParams struct {

	/*Body
	  Datagroup

	*/
	Body *models.Datagroup
	/*DatagroupID
	  ID of datagroup.

	*/
	DatagroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update datagroup params
func (o *UpdateDatagroupParams) WithTimeout(timeout time.Duration) *UpdateDatagroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update datagroup params
func (o *UpdateDatagroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update datagroup params
func (o *UpdateDatagroupParams) WithContext(ctx context.Context) *UpdateDatagroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update datagroup params
func (o *UpdateDatagroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update datagroup params
func (o *UpdateDatagroupParams) WithHTTPClient(client *http.Client) *UpdateDatagroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update datagroup params
func (o *UpdateDatagroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update datagroup params
func (o *UpdateDatagroupParams) WithBody(body *models.Datagroup) *UpdateDatagroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update datagroup params
func (o *UpdateDatagroupParams) SetBody(body *models.Datagroup) {
	o.Body = body
}

// WithDatagroupID adds the datagroupID to the update datagroup params
func (o *UpdateDatagroupParams) WithDatagroupID(datagroupID string) *UpdateDatagroupParams {
	o.SetDatagroupID(datagroupID)
	return o
}

// SetDatagroupID adds the datagroupId to the update datagroup params
func (o *UpdateDatagroupParams) SetDatagroupID(datagroupID string) {
	o.DatagroupID = datagroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDatagroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param datagroup_id
	if err := r.SetPathParam("datagroup_id", o.DatagroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
