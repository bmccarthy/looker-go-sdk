// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// AllProjectsReader is a Reader for the AllProjects structure.
type AllProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllProjectsOK creates a AllProjectsOK with default headers values
func NewAllProjectsOK() *AllProjectsOK {
	return &AllProjectsOK{}
}

/*AllProjectsOK handles this case with default header values.

Project
*/
type AllProjectsOK struct {
	Payload []*models.Project
}

func (o *AllProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] allProjectsOK  %+v", 200, o.Payload)
}

func (o *AllProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllProjectsBadRequest creates a AllProjectsBadRequest with default headers values
func NewAllProjectsBadRequest() *AllProjectsBadRequest {
	return &AllProjectsBadRequest{}
}

/*AllProjectsBadRequest handles this case with default header values.

Bad Request
*/
type AllProjectsBadRequest struct {
	Payload *models.Error
}

func (o *AllProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects][%d] allProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *AllProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllProjectsNotFound creates a AllProjectsNotFound with default headers values
func NewAllProjectsNotFound() *AllProjectsNotFound {
	return &AllProjectsNotFound{}
}

/*AllProjectsNotFound handles this case with default header values.

Not Found
*/
type AllProjectsNotFound struct {
	Payload *models.Error
}

func (o *AllProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects][%d] allProjectsNotFound  %+v", 404, o.Payload)
}

func (o *AllProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
