// Code generated by go-swagger; DO NOT EDIT.

package render_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateLookRenderTaskReader is a Reader for the CreateLookRenderTask structure.
type CreateLookRenderTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLookRenderTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLookRenderTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLookRenderTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLookRenderTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateLookRenderTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateLookRenderTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLookRenderTaskOK creates a CreateLookRenderTaskOK with default headers values
func NewCreateLookRenderTaskOK() *CreateLookRenderTaskOK {
	return &CreateLookRenderTaskOK{}
}

/*CreateLookRenderTaskOK handles this case with default header values.

Render Task
*/
type CreateLookRenderTaskOK struct {
	Payload *models.RenderTask
}

func (o *CreateLookRenderTaskOK) Error() string {
	return fmt.Sprintf("[POST /render_tasks/looks/{look_id}/{result_format}][%d] createLookRenderTaskOK  %+v", 200, o.Payload)
}

func (o *CreateLookRenderTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RenderTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookRenderTaskBadRequest creates a CreateLookRenderTaskBadRequest with default headers values
func NewCreateLookRenderTaskBadRequest() *CreateLookRenderTaskBadRequest {
	return &CreateLookRenderTaskBadRequest{}
}

/*CreateLookRenderTaskBadRequest handles this case with default header values.

Bad Request
*/
type CreateLookRenderTaskBadRequest struct {
	Payload *models.Error
}

func (o *CreateLookRenderTaskBadRequest) Error() string {
	return fmt.Sprintf("[POST /render_tasks/looks/{look_id}/{result_format}][%d] createLookRenderTaskBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLookRenderTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookRenderTaskNotFound creates a CreateLookRenderTaskNotFound with default headers values
func NewCreateLookRenderTaskNotFound() *CreateLookRenderTaskNotFound {
	return &CreateLookRenderTaskNotFound{}
}

/*CreateLookRenderTaskNotFound handles this case with default header values.

Not Found
*/
type CreateLookRenderTaskNotFound struct {
	Payload *models.Error
}

func (o *CreateLookRenderTaskNotFound) Error() string {
	return fmt.Sprintf("[POST /render_tasks/looks/{look_id}/{result_format}][%d] createLookRenderTaskNotFound  %+v", 404, o.Payload)
}

func (o *CreateLookRenderTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookRenderTaskConflict creates a CreateLookRenderTaskConflict with default headers values
func NewCreateLookRenderTaskConflict() *CreateLookRenderTaskConflict {
	return &CreateLookRenderTaskConflict{}
}

/*CreateLookRenderTaskConflict handles this case with default header values.

Resource Already Exists
*/
type CreateLookRenderTaskConflict struct {
	Payload *models.Error
}

func (o *CreateLookRenderTaskConflict) Error() string {
	return fmt.Sprintf("[POST /render_tasks/looks/{look_id}/{result_format}][%d] createLookRenderTaskConflict  %+v", 409, o.Payload)
}

func (o *CreateLookRenderTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookRenderTaskUnprocessableEntity creates a CreateLookRenderTaskUnprocessableEntity with default headers values
func NewCreateLookRenderTaskUnprocessableEntity() *CreateLookRenderTaskUnprocessableEntity {
	return &CreateLookRenderTaskUnprocessableEntity{}
}

/*CreateLookRenderTaskUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateLookRenderTaskUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateLookRenderTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /render_tasks/looks/{look_id}/{result_format}][%d] createLookRenderTaskUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateLookRenderTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}