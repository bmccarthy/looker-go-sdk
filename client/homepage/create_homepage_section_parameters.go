// Code generated by go-swagger; DO NOT EDIT.

package homepage

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

// NewCreateHomepageSectionParams creates a new CreateHomepageSectionParams object
// with the default values initialized.
func NewCreateHomepageSectionParams() *CreateHomepageSectionParams {
	var ()
	return &CreateHomepageSectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHomepageSectionParamsWithTimeout creates a new CreateHomepageSectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHomepageSectionParamsWithTimeout(timeout time.Duration) *CreateHomepageSectionParams {
	var ()
	return &CreateHomepageSectionParams{

		timeout: timeout,
	}
}

// NewCreateHomepageSectionParamsWithContext creates a new CreateHomepageSectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHomepageSectionParamsWithContext(ctx context.Context) *CreateHomepageSectionParams {
	var ()
	return &CreateHomepageSectionParams{

		Context: ctx,
	}
}

// NewCreateHomepageSectionParamsWithHTTPClient creates a new CreateHomepageSectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHomepageSectionParamsWithHTTPClient(client *http.Client) *CreateHomepageSectionParams {
	var ()
	return &CreateHomepageSectionParams{
		HTTPClient: client,
	}
}

/*CreateHomepageSectionParams contains all the parameters to send to the API endpoint
for the create homepage section operation typically these are written to a http.Request
*/
type CreateHomepageSectionParams struct {

	/*Body
	  Homepage section

	*/
	Body *models.HomepageSection
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create homepage section params
func (o *CreateHomepageSectionParams) WithTimeout(timeout time.Duration) *CreateHomepageSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create homepage section params
func (o *CreateHomepageSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create homepage section params
func (o *CreateHomepageSectionParams) WithContext(ctx context.Context) *CreateHomepageSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create homepage section params
func (o *CreateHomepageSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create homepage section params
func (o *CreateHomepageSectionParams) WithHTTPClient(client *http.Client) *CreateHomepageSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create homepage section params
func (o *CreateHomepageSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create homepage section params
func (o *CreateHomepageSectionParams) WithBody(body *models.HomepageSection) *CreateHomepageSectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create homepage section params
func (o *CreateHomepageSectionParams) SetBody(body *models.HomepageSection) {
	o.Body = body
}

// WithFields adds the fields to the create homepage section params
func (o *CreateHomepageSectionParams) WithFields(fields *string) *CreateHomepageSectionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create homepage section params
func (o *CreateHomepageSectionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHomepageSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
