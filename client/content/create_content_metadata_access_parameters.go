// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewCreateContentMetadataAccessParams creates a new CreateContentMetadataAccessParams object
// with the default values initialized.
func NewCreateContentMetadataAccessParams() *CreateContentMetadataAccessParams {
	var ()
	return &CreateContentMetadataAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateContentMetadataAccessParamsWithTimeout creates a new CreateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateContentMetadataAccessParamsWithTimeout(timeout time.Duration) *CreateContentMetadataAccessParams {
	var ()
	return &CreateContentMetadataAccessParams{

		timeout: timeout,
	}
}

// NewCreateContentMetadataAccessParamsWithContext creates a new CreateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateContentMetadataAccessParamsWithContext(ctx context.Context) *CreateContentMetadataAccessParams {
	var ()
	return &CreateContentMetadataAccessParams{

		Context: ctx,
	}
}

// NewCreateContentMetadataAccessParamsWithHTTPClient creates a new CreateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateContentMetadataAccessParamsWithHTTPClient(client *http.Client) *CreateContentMetadataAccessParams {
	var ()
	return &CreateContentMetadataAccessParams{
		HTTPClient: client,
	}
}

/*CreateContentMetadataAccessParams contains all the parameters to send to the API endpoint
for the create content metadata access operation typically these are written to a http.Request
*/
type CreateContentMetadataAccessParams struct {

	/*Body
	  Content Metadata Access

	*/
	Body *models.ContentMetaGroupUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create content metadata access params
func (o *CreateContentMetadataAccessParams) WithTimeout(timeout time.Duration) *CreateContentMetadataAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create content metadata access params
func (o *CreateContentMetadataAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create content metadata access params
func (o *CreateContentMetadataAccessParams) WithContext(ctx context.Context) *CreateContentMetadataAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create content metadata access params
func (o *CreateContentMetadataAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create content metadata access params
func (o *CreateContentMetadataAccessParams) WithHTTPClient(client *http.Client) *CreateContentMetadataAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create content metadata access params
func (o *CreateContentMetadataAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create content metadata access params
func (o *CreateContentMetadataAccessParams) WithBody(body *models.ContentMetaGroupUser) *CreateContentMetadataAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create content metadata access params
func (o *CreateContentMetadataAccessParams) SetBody(body *models.ContentMetaGroupUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateContentMetadataAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
