// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteGroupNoContent creates a DeleteGroupNoContent with default headers values
func NewDeleteGroupNoContent() *DeleteGroupNoContent {
	return &DeleteGroupNoContent{}
}

/*DeleteGroupNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteGroupNoContent struct {
	Payload string
}

func (o *DeleteGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}][%d] deleteGroupNoContent  %+v", 204, o.Payload)
}

func (o *DeleteGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupBadRequest creates a DeleteGroupBadRequest with default headers values
func NewDeleteGroupBadRequest() *DeleteGroupBadRequest {
	return &DeleteGroupBadRequest{}
}

/*DeleteGroupBadRequest handles this case with default header values.

Bad Request
*/
type DeleteGroupBadRequest struct {
	Payload *models.Error
}

func (o *DeleteGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}][%d] deleteGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupForbidden creates a DeleteGroupForbidden with default headers values
func NewDeleteGroupForbidden() *DeleteGroupForbidden {
	return &DeleteGroupForbidden{}
}

/*DeleteGroupForbidden handles this case with default header values.

Permission Denied
*/
type DeleteGroupForbidden struct {
	Payload *models.Error
}

func (o *DeleteGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}][%d] deleteGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupNotFound creates a DeleteGroupNotFound with default headers values
func NewDeleteGroupNotFound() *DeleteGroupNotFound {
	return &DeleteGroupNotFound{}
}

/*DeleteGroupNotFound handles this case with default header values.

Not Found
*/
type DeleteGroupNotFound struct {
	Payload *models.Error
}

func (o *DeleteGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}][%d] deleteGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
