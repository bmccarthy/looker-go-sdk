// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// DeleteUserAttributeUserValueReader is a Reader for the DeleteUserAttributeUserValue structure.
type DeleteUserAttributeUserValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserAttributeUserValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserAttributeUserValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserAttributeUserValueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserAttributeUserValueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserAttributeUserValueNoContent creates a DeleteUserAttributeUserValueNoContent with default headers values
func NewDeleteUserAttributeUserValueNoContent() *DeleteUserAttributeUserValueNoContent {
	return &DeleteUserAttributeUserValueNoContent{}
}

/*DeleteUserAttributeUserValueNoContent handles this case with default header values.

Deleted
*/
type DeleteUserAttributeUserValueNoContent struct {
}

func (o *DeleteUserAttributeUserValueNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/attribute_values/{user_attribute_id}][%d] deleteUserAttributeUserValueNoContent ", 204)
}

func (o *DeleteUserAttributeUserValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserAttributeUserValueBadRequest creates a DeleteUserAttributeUserValueBadRequest with default headers values
func NewDeleteUserAttributeUserValueBadRequest() *DeleteUserAttributeUserValueBadRequest {
	return &DeleteUserAttributeUserValueBadRequest{}
}

/*DeleteUserAttributeUserValueBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserAttributeUserValueBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserAttributeUserValueBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/attribute_values/{user_attribute_id}][%d] deleteUserAttributeUserValueBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserAttributeUserValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAttributeUserValueNotFound creates a DeleteUserAttributeUserValueNotFound with default headers values
func NewDeleteUserAttributeUserValueNotFound() *DeleteUserAttributeUserValueNotFound {
	return &DeleteUserAttributeUserValueNotFound{}
}

/*DeleteUserAttributeUserValueNotFound handles this case with default header values.

Not Found
*/
type DeleteUserAttributeUserValueNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserAttributeUserValueNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/attribute_values/{user_attribute_id}][%d] deleteUserAttributeUserValueNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserAttributeUserValueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
