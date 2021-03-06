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

// GroupReader is a Reader for the Group structure.
type GroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupOK creates a GroupOK with default headers values
func NewGroupOK() *GroupOK {
	return &GroupOK{}
}

/*GroupOK handles this case with default header values.

Group
*/
type GroupOK struct {
	Payload *models.Group
}

func (o *GroupOK) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}][%d] groupOK  %+v", 200, o.Payload)
}

func (o *GroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupBadRequest creates a GroupBadRequest with default headers values
func NewGroupBadRequest() *GroupBadRequest {
	return &GroupBadRequest{}
}

/*GroupBadRequest handles this case with default header values.

Bad Request
*/
type GroupBadRequest struct {
	Payload *models.Error
}

func (o *GroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}][%d] groupBadRequest  %+v", 400, o.Payload)
}

func (o *GroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupNotFound creates a GroupNotFound with default headers values
func NewGroupNotFound() *GroupNotFound {
	return &GroupNotFound{}
}

/*GroupNotFound handles this case with default header values.

Not Found
*/
type GroupNotFound struct {
	Payload *models.Error
}

func (o *GroupNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}][%d] groupNotFound  %+v", 404, o.Payload)
}

func (o *GroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
