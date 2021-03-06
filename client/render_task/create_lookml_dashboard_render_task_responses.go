// Code generated by go-swagger; DO NOT EDIT.

package render_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// CreateLookmlDashboardRenderTaskReader is a Reader for the CreateLookmlDashboardRenderTask structure.
type CreateLookmlDashboardRenderTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLookmlDashboardRenderTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLookmlDashboardRenderTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLookmlDashboardRenderTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLookmlDashboardRenderTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateLookmlDashboardRenderTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateLookmlDashboardRenderTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLookmlDashboardRenderTaskOK creates a CreateLookmlDashboardRenderTaskOK with default headers values
func NewCreateLookmlDashboardRenderTaskOK() *CreateLookmlDashboardRenderTaskOK {
	return &CreateLookmlDashboardRenderTaskOK{}
}

/*CreateLookmlDashboardRenderTaskOK handles this case with default header values.

Render Task
*/
type CreateLookmlDashboardRenderTaskOK struct {
	Payload *models.RenderTask
}

func (o *CreateLookmlDashboardRenderTaskOK) Error() string {
	return fmt.Sprintf("[POST /render_tasks/lookml_dashboards/{dashboard_id}/{result_format}][%d] createLookmlDashboardRenderTaskOK  %+v", 200, o.Payload)
}

func (o *CreateLookmlDashboardRenderTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RenderTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlDashboardRenderTaskBadRequest creates a CreateLookmlDashboardRenderTaskBadRequest with default headers values
func NewCreateLookmlDashboardRenderTaskBadRequest() *CreateLookmlDashboardRenderTaskBadRequest {
	return &CreateLookmlDashboardRenderTaskBadRequest{}
}

/*CreateLookmlDashboardRenderTaskBadRequest handles this case with default header values.

Bad Request
*/
type CreateLookmlDashboardRenderTaskBadRequest struct {
	Payload *models.Error
}

func (o *CreateLookmlDashboardRenderTaskBadRequest) Error() string {
	return fmt.Sprintf("[POST /render_tasks/lookml_dashboards/{dashboard_id}/{result_format}][%d] createLookmlDashboardRenderTaskBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLookmlDashboardRenderTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlDashboardRenderTaskNotFound creates a CreateLookmlDashboardRenderTaskNotFound with default headers values
func NewCreateLookmlDashboardRenderTaskNotFound() *CreateLookmlDashboardRenderTaskNotFound {
	return &CreateLookmlDashboardRenderTaskNotFound{}
}

/*CreateLookmlDashboardRenderTaskNotFound handles this case with default header values.

Not Found
*/
type CreateLookmlDashboardRenderTaskNotFound struct {
	Payload *models.Error
}

func (o *CreateLookmlDashboardRenderTaskNotFound) Error() string {
	return fmt.Sprintf("[POST /render_tasks/lookml_dashboards/{dashboard_id}/{result_format}][%d] createLookmlDashboardRenderTaskNotFound  %+v", 404, o.Payload)
}

func (o *CreateLookmlDashboardRenderTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlDashboardRenderTaskConflict creates a CreateLookmlDashboardRenderTaskConflict with default headers values
func NewCreateLookmlDashboardRenderTaskConflict() *CreateLookmlDashboardRenderTaskConflict {
	return &CreateLookmlDashboardRenderTaskConflict{}
}

/*CreateLookmlDashboardRenderTaskConflict handles this case with default header values.

Resource Already Exists
*/
type CreateLookmlDashboardRenderTaskConflict struct {
	Payload *models.Error
}

func (o *CreateLookmlDashboardRenderTaskConflict) Error() string {
	return fmt.Sprintf("[POST /render_tasks/lookml_dashboards/{dashboard_id}/{result_format}][%d] createLookmlDashboardRenderTaskConflict  %+v", 409, o.Payload)
}

func (o *CreateLookmlDashboardRenderTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlDashboardRenderTaskUnprocessableEntity creates a CreateLookmlDashboardRenderTaskUnprocessableEntity with default headers values
func NewCreateLookmlDashboardRenderTaskUnprocessableEntity() *CreateLookmlDashboardRenderTaskUnprocessableEntity {
	return &CreateLookmlDashboardRenderTaskUnprocessableEntity{}
}

/*CreateLookmlDashboardRenderTaskUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateLookmlDashboardRenderTaskUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateLookmlDashboardRenderTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /render_tasks/lookml_dashboards/{dashboard_id}/{result_format}][%d] createLookmlDashboardRenderTaskUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateLookmlDashboardRenderTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
