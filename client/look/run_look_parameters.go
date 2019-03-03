// Code generated by go-swagger; DO NOT EDIT.

package look

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

// NewRunLookParams creates a new RunLookParams object
// with the default values initialized.
func NewRunLookParams() *RunLookParams {
	var ()
	return &RunLookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunLookParamsWithTimeout creates a new RunLookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunLookParamsWithTimeout(timeout time.Duration) *RunLookParams {
	var ()
	return &RunLookParams{

		timeout: timeout,
	}
}

// NewRunLookParamsWithContext creates a new RunLookParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunLookParamsWithContext(ctx context.Context) *RunLookParams {
	var ()
	return &RunLookParams{

		Context: ctx,
	}
}

// NewRunLookParamsWithHTTPClient creates a new RunLookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunLookParamsWithHTTPClient(client *http.Client) *RunLookParams {
	var ()
	return &RunLookParams{
		HTTPClient: client,
	}
}

/*RunLookParams contains all the parameters to send to the API endpoint
for the run look operation typically these are written to a http.Request
*/
type RunLookParams struct {

	/*ApplyFormatting
	  Apply model-specified formatting to each result.

	*/
	ApplyFormatting *bool
	/*ApplyVis
	  Apply visualization options to results.

	*/
	ApplyVis *bool
	/*Cache
	  Get results from cache if available.

	*/
	Cache *bool
	/*CacheOnly
	  Retrieve any results from cache even if the results have expired.

	*/
	CacheOnly *bool
	/*ForceProduction
	  Force use of production models even if the user is in development mode.

	*/
	ForceProduction *bool
	/*GenerateDrillLinks
	  Generate drill links (only applicable to 'json_detail' format.

	*/
	GenerateDrillLinks *bool
	/*ImageHeight
	  Render height for image formats.

	*/
	ImageHeight *int64
	/*ImageWidth
	  Render width for image formats.

	*/
	ImageWidth *int64
	/*Limit
	  Row limit (may override the limit in the saved query).

	*/
	Limit *int64
	/*LookID
	  Id of look

	*/
	LookID int64
	/*PathPrefix
	  Prefix to use for drill links (url encoded).

	*/
	PathPrefix *string
	/*RebuildPdts
	  Rebuild PDTS used in query.

	*/
	RebuildPdts *bool
	/*ResultFormat
	  Format of result

	*/
	ResultFormat string
	/*ServerTableCalcs
	  Perform table calculations on query results

	*/
	ServerTableCalcs *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the run look params
func (o *RunLookParams) WithTimeout(timeout time.Duration) *RunLookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run look params
func (o *RunLookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run look params
func (o *RunLookParams) WithContext(ctx context.Context) *RunLookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run look params
func (o *RunLookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run look params
func (o *RunLookParams) WithHTTPClient(client *http.Client) *RunLookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run look params
func (o *RunLookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplyFormatting adds the applyFormatting to the run look params
func (o *RunLookParams) WithApplyFormatting(applyFormatting *bool) *RunLookParams {
	o.SetApplyFormatting(applyFormatting)
	return o
}

// SetApplyFormatting adds the applyFormatting to the run look params
func (o *RunLookParams) SetApplyFormatting(applyFormatting *bool) {
	o.ApplyFormatting = applyFormatting
}

// WithApplyVis adds the applyVis to the run look params
func (o *RunLookParams) WithApplyVis(applyVis *bool) *RunLookParams {
	o.SetApplyVis(applyVis)
	return o
}

// SetApplyVis adds the applyVis to the run look params
func (o *RunLookParams) SetApplyVis(applyVis *bool) {
	o.ApplyVis = applyVis
}

// WithCache adds the cache to the run look params
func (o *RunLookParams) WithCache(cache *bool) *RunLookParams {
	o.SetCache(cache)
	return o
}

// SetCache adds the cache to the run look params
func (o *RunLookParams) SetCache(cache *bool) {
	o.Cache = cache
}

// WithCacheOnly adds the cacheOnly to the run look params
func (o *RunLookParams) WithCacheOnly(cacheOnly *bool) *RunLookParams {
	o.SetCacheOnly(cacheOnly)
	return o
}

// SetCacheOnly adds the cacheOnly to the run look params
func (o *RunLookParams) SetCacheOnly(cacheOnly *bool) {
	o.CacheOnly = cacheOnly
}

// WithForceProduction adds the forceProduction to the run look params
func (o *RunLookParams) WithForceProduction(forceProduction *bool) *RunLookParams {
	o.SetForceProduction(forceProduction)
	return o
}

// SetForceProduction adds the forceProduction to the run look params
func (o *RunLookParams) SetForceProduction(forceProduction *bool) {
	o.ForceProduction = forceProduction
}

// WithGenerateDrillLinks adds the generateDrillLinks to the run look params
func (o *RunLookParams) WithGenerateDrillLinks(generateDrillLinks *bool) *RunLookParams {
	o.SetGenerateDrillLinks(generateDrillLinks)
	return o
}

// SetGenerateDrillLinks adds the generateDrillLinks to the run look params
func (o *RunLookParams) SetGenerateDrillLinks(generateDrillLinks *bool) {
	o.GenerateDrillLinks = generateDrillLinks
}

// WithImageHeight adds the imageHeight to the run look params
func (o *RunLookParams) WithImageHeight(imageHeight *int64) *RunLookParams {
	o.SetImageHeight(imageHeight)
	return o
}

// SetImageHeight adds the imageHeight to the run look params
func (o *RunLookParams) SetImageHeight(imageHeight *int64) {
	o.ImageHeight = imageHeight
}

// WithImageWidth adds the imageWidth to the run look params
func (o *RunLookParams) WithImageWidth(imageWidth *int64) *RunLookParams {
	o.SetImageWidth(imageWidth)
	return o
}

// SetImageWidth adds the imageWidth to the run look params
func (o *RunLookParams) SetImageWidth(imageWidth *int64) {
	o.ImageWidth = imageWidth
}

// WithLimit adds the limit to the run look params
func (o *RunLookParams) WithLimit(limit *int64) *RunLookParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the run look params
func (o *RunLookParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLookID adds the lookID to the run look params
func (o *RunLookParams) WithLookID(lookID int64) *RunLookParams {
	o.SetLookID(lookID)
	return o
}

// SetLookID adds the lookId to the run look params
func (o *RunLookParams) SetLookID(lookID int64) {
	o.LookID = lookID
}

// WithPathPrefix adds the pathPrefix to the run look params
func (o *RunLookParams) WithPathPrefix(pathPrefix *string) *RunLookParams {
	o.SetPathPrefix(pathPrefix)
	return o
}

// SetPathPrefix adds the pathPrefix to the run look params
func (o *RunLookParams) SetPathPrefix(pathPrefix *string) {
	o.PathPrefix = pathPrefix
}

// WithRebuildPdts adds the rebuildPdts to the run look params
func (o *RunLookParams) WithRebuildPdts(rebuildPdts *bool) *RunLookParams {
	o.SetRebuildPdts(rebuildPdts)
	return o
}

// SetRebuildPdts adds the rebuildPdts to the run look params
func (o *RunLookParams) SetRebuildPdts(rebuildPdts *bool) {
	o.RebuildPdts = rebuildPdts
}

// WithResultFormat adds the resultFormat to the run look params
func (o *RunLookParams) WithResultFormat(resultFormat string) *RunLookParams {
	o.SetResultFormat(resultFormat)
	return o
}

// SetResultFormat adds the resultFormat to the run look params
func (o *RunLookParams) SetResultFormat(resultFormat string) {
	o.ResultFormat = resultFormat
}

// WithServerTableCalcs adds the serverTableCalcs to the run look params
func (o *RunLookParams) WithServerTableCalcs(serverTableCalcs *bool) *RunLookParams {
	o.SetServerTableCalcs(serverTableCalcs)
	return o
}

// SetServerTableCalcs adds the serverTableCalcs to the run look params
func (o *RunLookParams) SetServerTableCalcs(serverTableCalcs *bool) {
	o.ServerTableCalcs = serverTableCalcs
}

// WriteToRequest writes these params to a swagger request
func (o *RunLookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplyFormatting != nil {

		// query param apply_formatting
		var qrApplyFormatting bool
		if o.ApplyFormatting != nil {
			qrApplyFormatting = *o.ApplyFormatting
		}
		qApplyFormatting := swag.FormatBool(qrApplyFormatting)
		if qApplyFormatting != "" {
			if err := r.SetQueryParam("apply_formatting", qApplyFormatting); err != nil {
				return err
			}
		}

	}

	if o.ApplyVis != nil {

		// query param apply_vis
		var qrApplyVis bool
		if o.ApplyVis != nil {
			qrApplyVis = *o.ApplyVis
		}
		qApplyVis := swag.FormatBool(qrApplyVis)
		if qApplyVis != "" {
			if err := r.SetQueryParam("apply_vis", qApplyVis); err != nil {
				return err
			}
		}

	}

	if o.Cache != nil {

		// query param cache
		var qrCache bool
		if o.Cache != nil {
			qrCache = *o.Cache
		}
		qCache := swag.FormatBool(qrCache)
		if qCache != "" {
			if err := r.SetQueryParam("cache", qCache); err != nil {
				return err
			}
		}

	}

	if o.CacheOnly != nil {

		// query param cache_only
		var qrCacheOnly bool
		if o.CacheOnly != nil {
			qrCacheOnly = *o.CacheOnly
		}
		qCacheOnly := swag.FormatBool(qrCacheOnly)
		if qCacheOnly != "" {
			if err := r.SetQueryParam("cache_only", qCacheOnly); err != nil {
				return err
			}
		}

	}

	if o.ForceProduction != nil {

		// query param force_production
		var qrForceProduction bool
		if o.ForceProduction != nil {
			qrForceProduction = *o.ForceProduction
		}
		qForceProduction := swag.FormatBool(qrForceProduction)
		if qForceProduction != "" {
			if err := r.SetQueryParam("force_production", qForceProduction); err != nil {
				return err
			}
		}

	}

	if o.GenerateDrillLinks != nil {

		// query param generate_drill_links
		var qrGenerateDrillLinks bool
		if o.GenerateDrillLinks != nil {
			qrGenerateDrillLinks = *o.GenerateDrillLinks
		}
		qGenerateDrillLinks := swag.FormatBool(qrGenerateDrillLinks)
		if qGenerateDrillLinks != "" {
			if err := r.SetQueryParam("generate_drill_links", qGenerateDrillLinks); err != nil {
				return err
			}
		}

	}

	if o.ImageHeight != nil {

		// query param image_height
		var qrImageHeight int64
		if o.ImageHeight != nil {
			qrImageHeight = *o.ImageHeight
		}
		qImageHeight := swag.FormatInt64(qrImageHeight)
		if qImageHeight != "" {
			if err := r.SetQueryParam("image_height", qImageHeight); err != nil {
				return err
			}
		}

	}

	if o.ImageWidth != nil {

		// query param image_width
		var qrImageWidth int64
		if o.ImageWidth != nil {
			qrImageWidth = *o.ImageWidth
		}
		qImageWidth := swag.FormatInt64(qrImageWidth)
		if qImageWidth != "" {
			if err := r.SetQueryParam("image_width", qImageWidth); err != nil {
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

	// path param look_id
	if err := r.SetPathParam("look_id", swag.FormatInt64(o.LookID)); err != nil {
		return err
	}

	if o.PathPrefix != nil {

		// query param path_prefix
		var qrPathPrefix string
		if o.PathPrefix != nil {
			qrPathPrefix = *o.PathPrefix
		}
		qPathPrefix := qrPathPrefix
		if qPathPrefix != "" {
			if err := r.SetQueryParam("path_prefix", qPathPrefix); err != nil {
				return err
			}
		}

	}

	if o.RebuildPdts != nil {

		// query param rebuild_pdts
		var qrRebuildPdts bool
		if o.RebuildPdts != nil {
			qrRebuildPdts = *o.RebuildPdts
		}
		qRebuildPdts := swag.FormatBool(qrRebuildPdts)
		if qRebuildPdts != "" {
			if err := r.SetQueryParam("rebuild_pdts", qRebuildPdts); err != nil {
				return err
			}
		}

	}

	// path param result_format
	if err := r.SetPathParam("result_format", o.ResultFormat); err != nil {
		return err
	}

	if o.ServerTableCalcs != nil {

		// query param server_table_calcs
		var qrServerTableCalcs bool
		if o.ServerTableCalcs != nil {
			qrServerTableCalcs = *o.ServerTableCalcs
		}
		qServerTableCalcs := swag.FormatBool(qrServerTableCalcs)
		if qServerTableCalcs != "" {
			if err := r.SetQueryParam("server_table_calcs", qServerTableCalcs); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
