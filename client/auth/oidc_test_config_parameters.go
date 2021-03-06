// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewOidcTestConfigParams creates a new OidcTestConfigParams object
// with the default values initialized.
func NewOidcTestConfigParams() *OidcTestConfigParams {
	var ()
	return &OidcTestConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOidcTestConfigParamsWithTimeout creates a new OidcTestConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOidcTestConfigParamsWithTimeout(timeout time.Duration) *OidcTestConfigParams {
	var ()
	return &OidcTestConfigParams{

		timeout: timeout,
	}
}

// NewOidcTestConfigParamsWithContext creates a new OidcTestConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewOidcTestConfigParamsWithContext(ctx context.Context) *OidcTestConfigParams {
	var ()
	return &OidcTestConfigParams{

		Context: ctx,
	}
}

// NewOidcTestConfigParamsWithHTTPClient creates a new OidcTestConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOidcTestConfigParamsWithHTTPClient(client *http.Client) *OidcTestConfigParams {
	var ()
	return &OidcTestConfigParams{
		HTTPClient: client,
	}
}

/*OidcTestConfigParams contains all the parameters to send to the API endpoint
for the oidc test config operation typically these are written to a http.Request
*/
type OidcTestConfigParams struct {

	/*TestSlug
	  Slug of test config

	*/
	TestSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oidc test config params
func (o *OidcTestConfigParams) WithTimeout(timeout time.Duration) *OidcTestConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oidc test config params
func (o *OidcTestConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oidc test config params
func (o *OidcTestConfigParams) WithContext(ctx context.Context) *OidcTestConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oidc test config params
func (o *OidcTestConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oidc test config params
func (o *OidcTestConfigParams) WithHTTPClient(client *http.Client) *OidcTestConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oidc test config params
func (o *OidcTestConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTestSlug adds the testSlug to the oidc test config params
func (o *OidcTestConfigParams) WithTestSlug(testSlug string) *OidcTestConfigParams {
	o.SetTestSlug(testSlug)
	return o
}

// SetTestSlug adds the testSlug to the oidc test config params
func (o *OidcTestConfigParams) SetTestSlug(testSlug string) {
	o.TestSlug = testSlug
}

// WriteToRequest writes these params to a swagger request
func (o *OidcTestConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param test_slug
	if err := r.SetPathParam("test_slug", o.TestSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
