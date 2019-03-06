// Code generated by go-swagger; DO NOT EDIT.

package config

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
)

// NewWhitelabelConfigurationParams creates a new WhitelabelConfigurationParams object
// with the default values initialized.
func NewWhitelabelConfigurationParams() *WhitelabelConfigurationParams {
	var ()
	return &WhitelabelConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWhitelabelConfigurationParamsWithTimeout creates a new WhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWhitelabelConfigurationParamsWithTimeout(timeout time.Duration) *WhitelabelConfigurationParams {
	var ()
	return &WhitelabelConfigurationParams{

		timeout: timeout,
	}
}

// NewWhitelabelConfigurationParamsWithContext creates a new WhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewWhitelabelConfigurationParamsWithContext(ctx context.Context) *WhitelabelConfigurationParams {
	var ()
	return &WhitelabelConfigurationParams{

		Context: ctx,
	}
}

// NewWhitelabelConfigurationParamsWithHTTPClient creates a new WhitelabelConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWhitelabelConfigurationParamsWithHTTPClient(client *http.Client) *WhitelabelConfigurationParams {
	var ()
	return &WhitelabelConfigurationParams{
		HTTPClient: client,
	}
}

/*WhitelabelConfigurationParams contains all the parameters to send to the API endpoint
for the whitelabel configuration operation typically these are written to a http.Request
*/
type WhitelabelConfigurationParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) WithTimeout(timeout time.Duration) *WhitelabelConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) WithContext(ctx context.Context) *WhitelabelConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) WithHTTPClient(client *http.Client) *WhitelabelConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) WithFields(fields *string) *WhitelabelConfigurationParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the whitelabel configuration params
func (o *WhitelabelConfigurationParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *WhitelabelConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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