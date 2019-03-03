// Code generated by go-swagger; DO NOT EDIT.

package session

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

// NewUpdateSessionParams creates a new UpdateSessionParams object
// with the default values initialized.
func NewUpdateSessionParams() *UpdateSessionParams {
	var ()
	return &UpdateSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSessionParamsWithTimeout creates a new UpdateSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSessionParamsWithTimeout(timeout time.Duration) *UpdateSessionParams {
	var ()
	return &UpdateSessionParams{

		timeout: timeout,
	}
}

// NewUpdateSessionParamsWithContext creates a new UpdateSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSessionParamsWithContext(ctx context.Context) *UpdateSessionParams {
	var ()
	return &UpdateSessionParams{

		Context: ctx,
	}
}

// NewUpdateSessionParamsWithHTTPClient creates a new UpdateSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSessionParamsWithHTTPClient(client *http.Client) *UpdateSessionParams {
	var ()
	return &UpdateSessionParams{
		HTTPClient: client,
	}
}

/*UpdateSessionParams contains all the parameters to send to the API endpoint
for the update session operation typically these are written to a http.Request
*/
type UpdateSessionParams struct {

	/*Body
	  Session

	*/
	Body *models.APISession

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update session params
func (o *UpdateSessionParams) WithTimeout(timeout time.Duration) *UpdateSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update session params
func (o *UpdateSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update session params
func (o *UpdateSessionParams) WithContext(ctx context.Context) *UpdateSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update session params
func (o *UpdateSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update session params
func (o *UpdateSessionParams) WithHTTPClient(client *http.Client) *UpdateSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update session params
func (o *UpdateSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update session params
func (o *UpdateSessionParams) WithBody(body *models.APISession) *UpdateSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update session params
func (o *UpdateSessionParams) SetBody(body *models.APISession) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
