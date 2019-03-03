// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewCreateGitDeployKeyParams creates a new CreateGitDeployKeyParams object
// with the default values initialized.
func NewCreateGitDeployKeyParams() *CreateGitDeployKeyParams {
	var ()
	return &CreateGitDeployKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGitDeployKeyParamsWithTimeout creates a new CreateGitDeployKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGitDeployKeyParamsWithTimeout(timeout time.Duration) *CreateGitDeployKeyParams {
	var ()
	return &CreateGitDeployKeyParams{

		timeout: timeout,
	}
}

// NewCreateGitDeployKeyParamsWithContext creates a new CreateGitDeployKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGitDeployKeyParamsWithContext(ctx context.Context) *CreateGitDeployKeyParams {
	var ()
	return &CreateGitDeployKeyParams{

		Context: ctx,
	}
}

// NewCreateGitDeployKeyParamsWithHTTPClient creates a new CreateGitDeployKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGitDeployKeyParamsWithHTTPClient(client *http.Client) *CreateGitDeployKeyParams {
	var ()
	return &CreateGitDeployKeyParams{
		HTTPClient: client,
	}
}

/*CreateGitDeployKeyParams contains all the parameters to send to the API endpoint
for the create git deploy key operation typically these are written to a http.Request
*/
type CreateGitDeployKeyParams struct {

	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create git deploy key params
func (o *CreateGitDeployKeyParams) WithTimeout(timeout time.Duration) *CreateGitDeployKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create git deploy key params
func (o *CreateGitDeployKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create git deploy key params
func (o *CreateGitDeployKeyParams) WithContext(ctx context.Context) *CreateGitDeployKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create git deploy key params
func (o *CreateGitDeployKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create git deploy key params
func (o *CreateGitDeployKeyParams) WithHTTPClient(client *http.Client) *CreateGitDeployKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create git deploy key params
func (o *CreateGitDeployKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the create git deploy key params
func (o *CreateGitDeployKeyParams) WithProjectID(projectID string) *CreateGitDeployKeyParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create git deploy key params
func (o *CreateGitDeployKeyParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGitDeployKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
