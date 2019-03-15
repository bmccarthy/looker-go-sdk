// Code generated by go-swagger; DO NOT EDIT.

package space

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
)

// NewSpaceChildrenSearchParams creates a new SpaceChildrenSearchParams object
// with the default values initialized.
func NewSpaceChildrenSearchParams() *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceChildrenSearchParamsWithTimeout creates a new SpaceChildrenSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpaceChildrenSearchParamsWithTimeout(timeout time.Duration) *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{

		timeout: timeout,
	}
}

// NewSpaceChildrenSearchParamsWithContext creates a new SpaceChildrenSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpaceChildrenSearchParamsWithContext(ctx context.Context) *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{

		Context: ctx,
	}
}

// NewSpaceChildrenSearchParamsWithHTTPClient creates a new SpaceChildrenSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpaceChildrenSearchParamsWithHTTPClient(client *http.Client) *SpaceChildrenSearchParams {
	var ()
	return &SpaceChildrenSearchParams{
		HTTPClient: client,
	}
}

/*SpaceChildrenSearchParams contains all the parameters to send to the API endpoint
for the space children search operation typically these are written to a http.Request
*/
type SpaceChildrenSearchParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Name
	  Match Space name.

	*/
	Name *string
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string
	/*SpaceID
	  Id of space

	*/
	SpaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the space children search params
func (o *SpaceChildrenSearchParams) WithTimeout(timeout time.Duration) *SpaceChildrenSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the space children search params
func (o *SpaceChildrenSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the space children search params
func (o *SpaceChildrenSearchParams) WithContext(ctx context.Context) *SpaceChildrenSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the space children search params
func (o *SpaceChildrenSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the space children search params
func (o *SpaceChildrenSearchParams) WithHTTPClient(client *http.Client) *SpaceChildrenSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the space children search params
func (o *SpaceChildrenSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the space children search params
func (o *SpaceChildrenSearchParams) WithFields(fields *string) *SpaceChildrenSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the space children search params
func (o *SpaceChildrenSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the space children search params
func (o *SpaceChildrenSearchParams) WithName(name *string) *SpaceChildrenSearchParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the space children search params
func (o *SpaceChildrenSearchParams) SetName(name *string) {
	o.Name = name
}

// WithSorts adds the sorts to the space children search params
func (o *SpaceChildrenSearchParams) WithSorts(sorts *string) *SpaceChildrenSearchParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the space children search params
func (o *SpaceChildrenSearchParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithSpaceID adds the spaceID to the space children search params
func (o *SpaceChildrenSearchParams) WithSpaceID(spaceID int64) *SpaceChildrenSearchParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the space children search params
func (o *SpaceChildrenSearchParams) SetSpaceID(spaceID int64) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceChildrenSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string
		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {
			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}

	}

	// path param space_id
	if err := r.SetPathParam("space_id", swag.FormatInt64(o.SpaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
