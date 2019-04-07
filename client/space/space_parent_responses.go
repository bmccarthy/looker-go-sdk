// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// SpaceParentReader is a Reader for the SpaceParent structure.
type SpaceParentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceParentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpaceParentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpaceParentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpaceParentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSpaceParentOK creates a SpaceParentOK with default headers values
func NewSpaceParentOK() *SpaceParentOK {
	return &SpaceParentOK{}
}

/*SpaceParentOK handles this case with default header values.

Space
*/
type SpaceParentOK struct {
	Payload *models.Space
}

func (o *SpaceParentOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/parent][%d] spaceParentOK  %+v", 200, o.Payload)
}

func (o *SpaceParentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Space)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceParentBadRequest creates a SpaceParentBadRequest with default headers values
func NewSpaceParentBadRequest() *SpaceParentBadRequest {
	return &SpaceParentBadRequest{}
}

/*SpaceParentBadRequest handles this case with default header values.

Bad Request
*/
type SpaceParentBadRequest struct {
	Payload *models.Error
}

func (o *SpaceParentBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/parent][%d] spaceParentBadRequest  %+v", 400, o.Payload)
}

func (o *SpaceParentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceParentNotFound creates a SpaceParentNotFound with default headers values
func NewSpaceParentNotFound() *SpaceParentNotFound {
	return &SpaceParentNotFound{}
}

/*SpaceParentNotFound handles this case with default header values.

Not Found
*/
type SpaceParentNotFound struct {
	Payload *models.Error
}

func (o *SpaceParentNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/parent][%d] spaceParentNotFound  %+v", 404, o.Payload)
}

func (o *SpaceParentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
