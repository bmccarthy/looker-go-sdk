// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// UpdateSessionReader is a Reader for the UpdateSession structure.
type UpdateSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateSessionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSessionOK creates a UpdateSessionOK with default headers values
func NewUpdateSessionOK() *UpdateSessionOK {
	return &UpdateSessionOK{}
}

/*UpdateSessionOK handles this case with default header values.

Session
*/
type UpdateSessionOK struct {
	Payload *models.APISession
}

func (o *UpdateSessionOK) Error() string {
	return fmt.Sprintf("[PATCH /session][%d] updateSessionOK  %+v", 200, o.Payload)
}

func (o *UpdateSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionBadRequest creates a UpdateSessionBadRequest with default headers values
func NewUpdateSessionBadRequest() *UpdateSessionBadRequest {
	return &UpdateSessionBadRequest{}
}

/*UpdateSessionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSessionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSessionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /session][%d] updateSessionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionNotFound creates a UpdateSessionNotFound with default headers values
func NewUpdateSessionNotFound() *UpdateSessionNotFound {
	return &UpdateSessionNotFound{}
}

/*UpdateSessionNotFound handles this case with default header values.

Not Found
*/
type UpdateSessionNotFound struct {
	Payload *models.Error
}

func (o *UpdateSessionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /session][%d] updateSessionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionUnprocessableEntity creates a UpdateSessionUnprocessableEntity with default headers values
func NewUpdateSessionUnprocessableEntity() *UpdateSessionUnprocessableEntity {
	return &UpdateSessionUnprocessableEntity{}
}

/*UpdateSessionUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateSessionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateSessionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /session][%d] updateSessionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateSessionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
