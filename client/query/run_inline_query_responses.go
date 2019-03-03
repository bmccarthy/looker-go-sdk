// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// RunInlineQueryReader is a Reader for the RunInlineQuery structure.
type RunInlineQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunInlineQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunInlineQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRunInlineQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRunInlineQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunInlineQueryOK creates a RunInlineQueryOK with default headers values
func NewRunInlineQueryOK() *RunInlineQueryOK {
	return &RunInlineQueryOK{}
}

/*RunInlineQueryOK handles this case with default header values.

Query Result
*/
type RunInlineQueryOK struct {
	Payload string
}

func (o *RunInlineQueryOK) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryOK  %+v", 200, o.Payload)
}

func (o *RunInlineQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryBadRequest creates a RunInlineQueryBadRequest with default headers values
func NewRunInlineQueryBadRequest() *RunInlineQueryBadRequest {
	return &RunInlineQueryBadRequest{}
}

/*RunInlineQueryBadRequest handles this case with default header values.

Bad Request
*/
type RunInlineQueryBadRequest struct {
	Payload *models.Error
}

func (o *RunInlineQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryBadRequest  %+v", 400, o.Payload)
}

func (o *RunInlineQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryNotFound creates a RunInlineQueryNotFound with default headers values
func NewRunInlineQueryNotFound() *RunInlineQueryNotFound {
	return &RunInlineQueryNotFound{}
}

/*RunInlineQueryNotFound handles this case with default header values.

Not Found
*/
type RunInlineQueryNotFound struct {
	Payload *models.Error
}

func (o *RunInlineQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryNotFound  %+v", 404, o.Payload)
}

func (o *RunInlineQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
