// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

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

	models "github.com/billtrust/looker-go-sdk/models"
)

// NewUpdateScheduledPlanParams creates a new UpdateScheduledPlanParams object
// with the default values initialized.
func NewUpdateScheduledPlanParams() *UpdateScheduledPlanParams {
	var ()
	return &UpdateScheduledPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScheduledPlanParamsWithTimeout creates a new UpdateScheduledPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateScheduledPlanParamsWithTimeout(timeout time.Duration) *UpdateScheduledPlanParams {
	var ()
	return &UpdateScheduledPlanParams{

		timeout: timeout,
	}
}

// NewUpdateScheduledPlanParamsWithContext creates a new UpdateScheduledPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateScheduledPlanParamsWithContext(ctx context.Context) *UpdateScheduledPlanParams {
	var ()
	return &UpdateScheduledPlanParams{

		Context: ctx,
	}
}

// NewUpdateScheduledPlanParamsWithHTTPClient creates a new UpdateScheduledPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateScheduledPlanParamsWithHTTPClient(client *http.Client) *UpdateScheduledPlanParams {
	var ()
	return &UpdateScheduledPlanParams{
		HTTPClient: client,
	}
}

/*UpdateScheduledPlanParams contains all the parameters to send to the API endpoint
for the update scheduled plan operation typically these are written to a http.Request
*/
type UpdateScheduledPlanParams struct {

	/*Body
	  Scheduled Plan

	*/
	Body *models.ScheduledPlan
	/*ScheduledPlanID
	  Scheduled Plan Id

	*/
	ScheduledPlanID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update scheduled plan params
func (o *UpdateScheduledPlanParams) WithTimeout(timeout time.Duration) *UpdateScheduledPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scheduled plan params
func (o *UpdateScheduledPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scheduled plan params
func (o *UpdateScheduledPlanParams) WithContext(ctx context.Context) *UpdateScheduledPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scheduled plan params
func (o *UpdateScheduledPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scheduled plan params
func (o *UpdateScheduledPlanParams) WithHTTPClient(client *http.Client) *UpdateScheduledPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scheduled plan params
func (o *UpdateScheduledPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update scheduled plan params
func (o *UpdateScheduledPlanParams) WithBody(body *models.ScheduledPlan) *UpdateScheduledPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update scheduled plan params
func (o *UpdateScheduledPlanParams) SetBody(body *models.ScheduledPlan) {
	o.Body = body
}

// WithScheduledPlanID adds the scheduledPlanID to the update scheduled plan params
func (o *UpdateScheduledPlanParams) WithScheduledPlanID(scheduledPlanID int64) *UpdateScheduledPlanParams {
	o.SetScheduledPlanID(scheduledPlanID)
	return o
}

// SetScheduledPlanID adds the scheduledPlanId to the update scheduled plan params
func (o *UpdateScheduledPlanParams) SetScheduledPlanID(scheduledPlanID int64) {
	o.ScheduledPlanID = scheduledPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduledPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scheduled_plan_id
	if err := r.SetPathParam("scheduled_plan_id", swag.FormatInt64(o.ScheduledPlanID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
