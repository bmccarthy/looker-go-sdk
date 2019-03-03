// Code generated by go-swagger; DO NOT EDIT.

package render_task

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

	models "looker-go-sdk/models"
)

// NewCreateLookmlDashboardRenderTaskParams creates a new CreateLookmlDashboardRenderTaskParams object
// with the default values initialized.
func NewCreateLookmlDashboardRenderTaskParams() *CreateLookmlDashboardRenderTaskParams {
	var ()
	return &CreateLookmlDashboardRenderTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLookmlDashboardRenderTaskParamsWithTimeout creates a new CreateLookmlDashboardRenderTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLookmlDashboardRenderTaskParamsWithTimeout(timeout time.Duration) *CreateLookmlDashboardRenderTaskParams {
	var ()
	return &CreateLookmlDashboardRenderTaskParams{

		timeout: timeout,
	}
}

// NewCreateLookmlDashboardRenderTaskParamsWithContext creates a new CreateLookmlDashboardRenderTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLookmlDashboardRenderTaskParamsWithContext(ctx context.Context) *CreateLookmlDashboardRenderTaskParams {
	var ()
	return &CreateLookmlDashboardRenderTaskParams{

		Context: ctx,
	}
}

// NewCreateLookmlDashboardRenderTaskParamsWithHTTPClient creates a new CreateLookmlDashboardRenderTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLookmlDashboardRenderTaskParamsWithHTTPClient(client *http.Client) *CreateLookmlDashboardRenderTaskParams {
	var ()
	return &CreateLookmlDashboardRenderTaskParams{
		HTTPClient: client,
	}
}

/*CreateLookmlDashboardRenderTaskParams contains all the parameters to send to the API endpoint
for the create lookml dashboard render task operation typically these are written to a http.Request
*/
type CreateLookmlDashboardRenderTaskParams struct {

	/*Body
	  Dashboard render task parameters

	*/
	Body *models.CreateDashboardRenderTask
	/*DashboardID
	  Id of lookml dashboard to render

	*/
	DashboardID string
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Height
	  Output height in pixels

	*/
	Height int64
	/*PdfLandscape
	  Whether to render pdf in landscape

	*/
	PdfLandscape *bool
	/*PdfPaperSize
	  Paper size for pdf

	*/
	PdfPaperSize *string
	/*ResultFormat
	  Output type: pdf, png, or jpg

	*/
	ResultFormat string
	/*Width
	  Output width in pixels

	*/
	Width int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithTimeout(timeout time.Duration) *CreateLookmlDashboardRenderTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithContext(ctx context.Context) *CreateLookmlDashboardRenderTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithHTTPClient(client *http.Client) *CreateLookmlDashboardRenderTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithBody(body *models.CreateDashboardRenderTask) *CreateLookmlDashboardRenderTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetBody(body *models.CreateDashboardRenderTask) {
	o.Body = body
}

// WithDashboardID adds the dashboardID to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithDashboardID(dashboardID string) *CreateLookmlDashboardRenderTaskParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WithFields adds the fields to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithFields(fields *string) *CreateLookmlDashboardRenderTaskParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHeight adds the height to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithHeight(height int64) *CreateLookmlDashboardRenderTaskParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetHeight(height int64) {
	o.Height = height
}

// WithPdfLandscape adds the pdfLandscape to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithPdfLandscape(pdfLandscape *bool) *CreateLookmlDashboardRenderTaskParams {
	o.SetPdfLandscape(pdfLandscape)
	return o
}

// SetPdfLandscape adds the pdfLandscape to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetPdfLandscape(pdfLandscape *bool) {
	o.PdfLandscape = pdfLandscape
}

// WithPdfPaperSize adds the pdfPaperSize to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithPdfPaperSize(pdfPaperSize *string) *CreateLookmlDashboardRenderTaskParams {
	o.SetPdfPaperSize(pdfPaperSize)
	return o
}

// SetPdfPaperSize adds the pdfPaperSize to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetPdfPaperSize(pdfPaperSize *string) {
	o.PdfPaperSize = pdfPaperSize
}

// WithResultFormat adds the resultFormat to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithResultFormat(resultFormat string) *CreateLookmlDashboardRenderTaskParams {
	o.SetResultFormat(resultFormat)
	return o
}

// SetResultFormat adds the resultFormat to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetResultFormat(resultFormat string) {
	o.ResultFormat = resultFormat
}

// WithWidth adds the width to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) WithWidth(width int64) *CreateLookmlDashboardRenderTaskParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the create lookml dashboard render task params
func (o *CreateLookmlDashboardRenderTaskParams) SetWidth(width int64) {
	o.Width = width
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLookmlDashboardRenderTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboard_id
	if err := r.SetPathParam("dashboard_id", o.DashboardID); err != nil {
		return err
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

	// query param height
	qrHeight := o.Height
	qHeight := swag.FormatInt64(qrHeight)
	if qHeight != "" {
		if err := r.SetQueryParam("height", qHeight); err != nil {
			return err
		}
	}

	if o.PdfLandscape != nil {

		// query param pdf_landscape
		var qrPdfLandscape bool
		if o.PdfLandscape != nil {
			qrPdfLandscape = *o.PdfLandscape
		}
		qPdfLandscape := swag.FormatBool(qrPdfLandscape)
		if qPdfLandscape != "" {
			if err := r.SetQueryParam("pdf_landscape", qPdfLandscape); err != nil {
				return err
			}
		}

	}

	if o.PdfPaperSize != nil {

		// query param pdf_paper_size
		var qrPdfPaperSize string
		if o.PdfPaperSize != nil {
			qrPdfPaperSize = *o.PdfPaperSize
		}
		qPdfPaperSize := qrPdfPaperSize
		if qPdfPaperSize != "" {
			if err := r.SetQueryParam("pdf_paper_size", qPdfPaperSize); err != nil {
				return err
			}
		}

	}

	// path param result_format
	if err := r.SetPathParam("result_format", o.ResultFormat); err != nil {
		return err
	}

	// query param width
	qrWidth := o.Width
	qWidth := swag.FormatInt64(qrWidth)
	if qWidth != "" {
		if err := r.SetQueryParam("width", qWidth); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
