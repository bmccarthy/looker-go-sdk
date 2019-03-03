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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchContentFavoritesParams creates a new SearchContentFavoritesParams object
// with the default values initialized.
func NewSearchContentFavoritesParams() *SearchContentFavoritesParams {
	var ()
	return &SearchContentFavoritesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchContentFavoritesParamsWithTimeout creates a new SearchContentFavoritesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchContentFavoritesParamsWithTimeout(timeout time.Duration) *SearchContentFavoritesParams {
	var ()
	return &SearchContentFavoritesParams{

		timeout: timeout,
	}
}

// NewSearchContentFavoritesParamsWithContext creates a new SearchContentFavoritesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchContentFavoritesParamsWithContext(ctx context.Context) *SearchContentFavoritesParams {
	var ()
	return &SearchContentFavoritesParams{

		Context: ctx,
	}
}

// NewSearchContentFavoritesParamsWithHTTPClient creates a new SearchContentFavoritesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchContentFavoritesParamsWithHTTPClient(client *http.Client) *SearchContentFavoritesParams {
	var ()
	return &SearchContentFavoritesParams{
		HTTPClient: client,
	}
}

/*SearchContentFavoritesParams contains all the parameters to send to the API endpoint
for the search content favorites operation typically these are written to a http.Request
*/
type SearchContentFavoritesParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Limit
	  Number of results to return. (used with offset)

	*/
	Limit *int64
	/*Offset
	  Number of results to skip before returning any. (used with limit)

	*/
	Offset *int64
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string
	/*UserID
	  Match User Id

	*/
	UserID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search content favorites params
func (o *SearchContentFavoritesParams) WithTimeout(timeout time.Duration) *SearchContentFavoritesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search content favorites params
func (o *SearchContentFavoritesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search content favorites params
func (o *SearchContentFavoritesParams) WithContext(ctx context.Context) *SearchContentFavoritesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search content favorites params
func (o *SearchContentFavoritesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search content favorites params
func (o *SearchContentFavoritesParams) WithHTTPClient(client *http.Client) *SearchContentFavoritesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search content favorites params
func (o *SearchContentFavoritesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the search content favorites params
func (o *SearchContentFavoritesParams) WithFields(fields *string) *SearchContentFavoritesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search content favorites params
func (o *SearchContentFavoritesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the search content favorites params
func (o *SearchContentFavoritesParams) WithLimit(limit *int64) *SearchContentFavoritesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search content favorites params
func (o *SearchContentFavoritesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search content favorites params
func (o *SearchContentFavoritesParams) WithOffset(offset *int64) *SearchContentFavoritesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search content favorites params
func (o *SearchContentFavoritesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSorts adds the sorts to the search content favorites params
func (o *SearchContentFavoritesParams) WithSorts(sorts *string) *SearchContentFavoritesParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search content favorites params
func (o *SearchContentFavoritesParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithUserID adds the userID to the search content favorites params
func (o *SearchContentFavoritesParams) WithUserID(userID *int64) *SearchContentFavoritesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the search content favorites params
func (o *SearchContentFavoritesParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchContentFavoritesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
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

	if o.UserID != nil {

		// query param user_id
		var qrUserID int64
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatInt64(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
