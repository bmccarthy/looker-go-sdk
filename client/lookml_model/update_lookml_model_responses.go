// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// UpdateLookmlModelReader is a Reader for the UpdateLookmlModel structure.
type UpdateLookmlModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLookmlModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateLookmlModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateLookmlModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateLookmlModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateLookmlModelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateLookmlModelOK creates a UpdateLookmlModelOK with default headers values
func NewUpdateLookmlModelOK() *UpdateLookmlModelOK {
	return &UpdateLookmlModelOK{}
}

/*UpdateLookmlModelOK handles this case with default header values.

LookML Model
*/
type UpdateLookmlModelOK struct {
	Payload *models.LookmlModel
}

func (o *UpdateLookmlModelOK) Error() string {
	return fmt.Sprintf("[PATCH /lookml_models/{lookml_model_name}][%d] updateLookmlModelOK  %+v", 200, o.Payload)
}

func (o *UpdateLookmlModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookmlModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookmlModelBadRequest creates a UpdateLookmlModelBadRequest with default headers values
func NewUpdateLookmlModelBadRequest() *UpdateLookmlModelBadRequest {
	return &UpdateLookmlModelBadRequest{}
}

/*UpdateLookmlModelBadRequest handles this case with default header values.

Bad Request
*/
type UpdateLookmlModelBadRequest struct {
	Payload *models.Error
}

func (o *UpdateLookmlModelBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /lookml_models/{lookml_model_name}][%d] updateLookmlModelBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLookmlModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookmlModelNotFound creates a UpdateLookmlModelNotFound with default headers values
func NewUpdateLookmlModelNotFound() *UpdateLookmlModelNotFound {
	return &UpdateLookmlModelNotFound{}
}

/*UpdateLookmlModelNotFound handles this case with default header values.

Not Found
*/
type UpdateLookmlModelNotFound struct {
	Payload *models.Error
}

func (o *UpdateLookmlModelNotFound) Error() string {
	return fmt.Sprintf("[PATCH /lookml_models/{lookml_model_name}][%d] updateLookmlModelNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLookmlModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLookmlModelUnprocessableEntity creates a UpdateLookmlModelUnprocessableEntity with default headers values
func NewUpdateLookmlModelUnprocessableEntity() *UpdateLookmlModelUnprocessableEntity {
	return &UpdateLookmlModelUnprocessableEntity{}
}

/*UpdateLookmlModelUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateLookmlModelUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateLookmlModelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /lookml_models/{lookml_model_name}][%d] updateLookmlModelUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateLookmlModelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
