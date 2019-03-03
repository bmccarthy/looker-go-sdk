// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// UpdateIntegrationReader is a Reader for the UpdateIntegration structure.
type UpdateIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateIntegrationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateIntegrationOK creates a UpdateIntegrationOK with default headers values
func NewUpdateIntegrationOK() *UpdateIntegrationOK {
	return &UpdateIntegrationOK{}
}

/*UpdateIntegrationOK handles this case with default header values.

Integration
*/
type UpdateIntegrationOK struct {
	Payload *models.Integration
}

func (o *UpdateIntegrationOK) Error() string {
	return fmt.Sprintf("[PATCH /integrations/{integration_id}][%d] updateIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationBadRequest creates a UpdateIntegrationBadRequest with default headers values
func NewUpdateIntegrationBadRequest() *UpdateIntegrationBadRequest {
	return &UpdateIntegrationBadRequest{}
}

/*UpdateIntegrationBadRequest handles this case with default header values.

Bad Request
*/
type UpdateIntegrationBadRequest struct {
	Payload *models.Error
}

func (o *UpdateIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /integrations/{integration_id}][%d] updateIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationNotFound creates a UpdateIntegrationNotFound with default headers values
func NewUpdateIntegrationNotFound() *UpdateIntegrationNotFound {
	return &UpdateIntegrationNotFound{}
}

/*UpdateIntegrationNotFound handles this case with default header values.

Not Found
*/
type UpdateIntegrationNotFound struct {
	Payload *models.Error
}

func (o *UpdateIntegrationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /integrations/{integration_id}][%d] updateIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIntegrationUnprocessableEntity creates a UpdateIntegrationUnprocessableEntity with default headers values
func NewUpdateIntegrationUnprocessableEntity() *UpdateIntegrationUnprocessableEntity {
	return &UpdateIntegrationUnprocessableEntity{}
}

/*UpdateIntegrationUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateIntegrationUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateIntegrationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /integrations/{integration_id}][%d] updateIntegrationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateIntegrationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
