// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// AllGitConnectionTestsReader is a Reader for the AllGitConnectionTests structure.
type AllGitConnectionTestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllGitConnectionTestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllGitConnectionTestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllGitConnectionTestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllGitConnectionTestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllGitConnectionTestsOK creates a AllGitConnectionTestsOK with default headers values
func NewAllGitConnectionTestsOK() *AllGitConnectionTestsOK {
	return &AllGitConnectionTestsOK{}
}

/*AllGitConnectionTestsOK handles this case with default header values.

Git Connection Test
*/
type AllGitConnectionTestsOK struct {
	Payload []*models.GitConnectionTest
}

func (o *AllGitConnectionTestsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests][%d] allGitConnectionTestsOK  %+v", 200, o.Payload)
}

func (o *AllGitConnectionTestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGitConnectionTestsBadRequest creates a AllGitConnectionTestsBadRequest with default headers values
func NewAllGitConnectionTestsBadRequest() *AllGitConnectionTestsBadRequest {
	return &AllGitConnectionTestsBadRequest{}
}

/*AllGitConnectionTestsBadRequest handles this case with default header values.

Bad Request
*/
type AllGitConnectionTestsBadRequest struct {
	Payload *models.Error
}

func (o *AllGitConnectionTestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests][%d] allGitConnectionTestsBadRequest  %+v", 400, o.Payload)
}

func (o *AllGitConnectionTestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGitConnectionTestsNotFound creates a AllGitConnectionTestsNotFound with default headers values
func NewAllGitConnectionTestsNotFound() *AllGitConnectionTestsNotFound {
	return &AllGitConnectionTestsNotFound{}
}

/*AllGitConnectionTestsNotFound handles this case with default header values.

Not Found
*/
type AllGitConnectionTestsNotFound struct {
	Payload *models.Error
}

func (o *AllGitConnectionTestsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests][%d] allGitConnectionTestsNotFound  %+v", 404, o.Payload)
}

func (o *AllGitConnectionTestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
