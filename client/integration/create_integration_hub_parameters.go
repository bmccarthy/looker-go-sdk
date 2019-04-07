// Code generated by go-swagger; DO NOT EDIT.

package integration

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

	models "github.com/billtrust/looker-go-sdk/models"
)

// NewCreateIntegrationHubParams creates a new CreateIntegrationHubParams object
// with the default values initialized.
func NewCreateIntegrationHubParams() *CreateIntegrationHubParams {
	var ()
	return &CreateIntegrationHubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIntegrationHubParamsWithTimeout creates a new CreateIntegrationHubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateIntegrationHubParamsWithTimeout(timeout time.Duration) *CreateIntegrationHubParams {
	var ()
	return &CreateIntegrationHubParams{

		timeout: timeout,
	}
}

// NewCreateIntegrationHubParamsWithContext creates a new CreateIntegrationHubParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateIntegrationHubParamsWithContext(ctx context.Context) *CreateIntegrationHubParams {
	var ()
	return &CreateIntegrationHubParams{

		Context: ctx,
	}
}

// NewCreateIntegrationHubParamsWithHTTPClient creates a new CreateIntegrationHubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateIntegrationHubParamsWithHTTPClient(client *http.Client) *CreateIntegrationHubParams {
	var ()
	return &CreateIntegrationHubParams{
		HTTPClient: client,
	}
}

/*CreateIntegrationHubParams contains all the parameters to send to the API endpoint
for the create integration hub operation typically these are written to a http.Request
*/
type CreateIntegrationHubParams struct {

	/*Body
	  Integration Hub

	*/
	Body *models.IntegrationHub
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create integration hub params
func (o *CreateIntegrationHubParams) WithTimeout(timeout time.Duration) *CreateIntegrationHubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create integration hub params
func (o *CreateIntegrationHubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create integration hub params
func (o *CreateIntegrationHubParams) WithContext(ctx context.Context) *CreateIntegrationHubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create integration hub params
func (o *CreateIntegrationHubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create integration hub params
func (o *CreateIntegrationHubParams) WithHTTPClient(client *http.Client) *CreateIntegrationHubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create integration hub params
func (o *CreateIntegrationHubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create integration hub params
func (o *CreateIntegrationHubParams) WithBody(body *models.IntegrationHub) *CreateIntegrationHubParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create integration hub params
func (o *CreateIntegrationHubParams) SetBody(body *models.IntegrationHub) {
	o.Body = body
}

// WithFields adds the fields to the create integration hub params
func (o *CreateIntegrationHubParams) WithFields(fields *string) *CreateIntegrationHubParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create integration hub params
func (o *CreateIntegrationHubParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIntegrationHubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
