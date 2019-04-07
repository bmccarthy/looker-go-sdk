// Code generated by go-swagger; DO NOT EDIT.

package connection

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

// NewCreateConnectionParams creates a new CreateConnectionParams object
// with the default values initialized.
func NewCreateConnectionParams() *CreateConnectionParams {
	var ()
	return &CreateConnectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateConnectionParamsWithTimeout creates a new CreateConnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateConnectionParamsWithTimeout(timeout time.Duration) *CreateConnectionParams {
	var ()
	return &CreateConnectionParams{

		timeout: timeout,
	}
}

// NewCreateConnectionParamsWithContext creates a new CreateConnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateConnectionParamsWithContext(ctx context.Context) *CreateConnectionParams {
	var ()
	return &CreateConnectionParams{

		Context: ctx,
	}
}

// NewCreateConnectionParamsWithHTTPClient creates a new CreateConnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateConnectionParamsWithHTTPClient(client *http.Client) *CreateConnectionParams {
	var ()
	return &CreateConnectionParams{
		HTTPClient: client,
	}
}

/*CreateConnectionParams contains all the parameters to send to the API endpoint
for the create connection operation typically these are written to a http.Request
*/
type CreateConnectionParams struct {

	/*Body
	  Connection

	*/
	Body *models.DBConnection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create connection params
func (o *CreateConnectionParams) WithTimeout(timeout time.Duration) *CreateConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create connection params
func (o *CreateConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create connection params
func (o *CreateConnectionParams) WithContext(ctx context.Context) *CreateConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create connection params
func (o *CreateConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create connection params
func (o *CreateConnectionParams) WithHTTPClient(client *http.Client) *CreateConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create connection params
func (o *CreateConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create connection params
func (o *CreateConnectionParams) WithBody(body *models.DBConnection) *CreateConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create connection params
func (o *CreateConnectionParams) SetBody(body *models.DBConnection) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
