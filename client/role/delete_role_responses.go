// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDeleteRoleMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoleNoContent creates a DeleteRoleNoContent with default headers values
func NewDeleteRoleNoContent() *DeleteRoleNoContent {
	return &DeleteRoleNoContent{}
}

/*DeleteRoleNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteRoleNoContent struct {
	Payload string
}

func (o *DeleteRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /roles/{role_id}][%d] deleteRoleNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleBadRequest creates a DeleteRoleBadRequest with default headers values
func NewDeleteRoleBadRequest() *DeleteRoleBadRequest {
	return &DeleteRoleBadRequest{}
}

/*DeleteRoleBadRequest handles this case with default header values.

Bad Request
*/
type DeleteRoleBadRequest struct {
	Payload *models.Error
}

func (o *DeleteRoleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /roles/{role_id}][%d] deleteRoleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleNotFound creates a DeleteRoleNotFound with default headers values
func NewDeleteRoleNotFound() *DeleteRoleNotFound {
	return &DeleteRoleNotFound{}
}

/*DeleteRoleNotFound handles this case with default header values.

Not Found
*/
type DeleteRoleNotFound struct {
	Payload *models.Error
}

func (o *DeleteRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /roles/{role_id}][%d] deleteRoleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleMethodNotAllowed creates a DeleteRoleMethodNotAllowed with default headers values
func NewDeleteRoleMethodNotAllowed() *DeleteRoleMethodNotAllowed {
	return &DeleteRoleMethodNotAllowed{}
}

/*DeleteRoleMethodNotAllowed handles this case with default header values.

Resource Can't Be Modified
*/
type DeleteRoleMethodNotAllowed struct {
	Payload *models.Error
}

func (o *DeleteRoleMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /roles/{role_id}][%d] deleteRoleMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteRoleMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
