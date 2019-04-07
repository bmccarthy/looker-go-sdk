// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// RunQueryReader is a Reader for the RunQuery structure.
type RunQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRunQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRunQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewRunQueryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunQueryOK creates a RunQueryOK with default headers values
func NewRunQueryOK() *RunQueryOK {
	return &RunQueryOK{}
}

/*RunQueryOK handles this case with default header values.

Query
*/
type RunQueryOK struct {
	Payload string
}

func (o *RunQueryOK) Error() string {
	return fmt.Sprintf("[GET /queries/{query_id}/run/{result_format}][%d] runQueryOK  %+v", 200, o.Payload)
}

func (o *RunQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunQueryBadRequest creates a RunQueryBadRequest with default headers values
func NewRunQueryBadRequest() *RunQueryBadRequest {
	return &RunQueryBadRequest{}
}

/*RunQueryBadRequest handles this case with default header values.

Bad Request
*/
type RunQueryBadRequest struct {
	Payload *models.Error
}

func (o *RunQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /queries/{query_id}/run/{result_format}][%d] runQueryBadRequest  %+v", 400, o.Payload)
}

func (o *RunQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunQueryNotFound creates a RunQueryNotFound with default headers values
func NewRunQueryNotFound() *RunQueryNotFound {
	return &RunQueryNotFound{}
}

/*RunQueryNotFound handles this case with default header values.

Not Found
*/
type RunQueryNotFound struct {
	Payload *models.Error
}

func (o *RunQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /queries/{query_id}/run/{result_format}][%d] runQueryNotFound  %+v", 404, o.Payload)
}

func (o *RunQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunQueryUnprocessableEntity creates a RunQueryUnprocessableEntity with default headers values
func NewRunQueryUnprocessableEntity() *RunQueryUnprocessableEntity {
	return &RunQueryUnprocessableEntity{}
}

/*RunQueryUnprocessableEntity handles this case with default header values.

Validation Error
*/
type RunQueryUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *RunQueryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /queries/{query_id}/run/{result_format}][%d] runQueryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RunQueryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
