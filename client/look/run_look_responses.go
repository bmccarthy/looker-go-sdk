// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// RunLookReader is a Reader for the RunLook structure.
type RunLookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunLookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunLookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRunLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRunLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewRunLookUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunLookOK creates a RunLookOK with default headers values
func NewRunLookOK() *RunLookOK {
	return &RunLookOK{}
}

/*RunLookOK handles this case with default header values.

Look
*/
type RunLookOK struct {
	Payload string
}

func (o *RunLookOK) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}/run/{result_format}][%d] runLookOK  %+v", 200, o.Payload)
}

func (o *RunLookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunLookBadRequest creates a RunLookBadRequest with default headers values
func NewRunLookBadRequest() *RunLookBadRequest {
	return &RunLookBadRequest{}
}

/*RunLookBadRequest handles this case with default header values.

Bad Request
*/
type RunLookBadRequest struct {
	Payload *models.Error
}

func (o *RunLookBadRequest) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}/run/{result_format}][%d] runLookBadRequest  %+v", 400, o.Payload)
}

func (o *RunLookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunLookNotFound creates a RunLookNotFound with default headers values
func NewRunLookNotFound() *RunLookNotFound {
	return &RunLookNotFound{}
}

/*RunLookNotFound handles this case with default header values.

Not Found
*/
type RunLookNotFound struct {
	Payload *models.Error
}

func (o *RunLookNotFound) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}/run/{result_format}][%d] runLookNotFound  %+v", 404, o.Payload)
}

func (o *RunLookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunLookUnprocessableEntity creates a RunLookUnprocessableEntity with default headers values
func NewRunLookUnprocessableEntity() *RunLookUnprocessableEntity {
	return &RunLookUnprocessableEntity{}
}

/*RunLookUnprocessableEntity handles this case with default header values.

Validation Error
*/
type RunLookUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *RunLookUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}/run/{result_format}][%d] runLookUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RunLookUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}