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

// CreateGitDeployKeyReader is a Reader for the CreateGitDeployKey structure.
type CreateGitDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGitDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGitDeployKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateGitDeployKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateGitDeployKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateGitDeployKeyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGitDeployKeyOK creates a CreateGitDeployKeyOK with default headers values
func NewCreateGitDeployKeyOK() *CreateGitDeployKeyOK {
	return &CreateGitDeployKeyOK{}
}

/*CreateGitDeployKeyOK handles this case with default header values.

Project
*/
type CreateGitDeployKeyOK struct {
	Payload string
}

func (o *CreateGitDeployKeyOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/git/deploy_key][%d] createGitDeployKeyOK  %+v", 200, o.Payload)
}

func (o *CreateGitDeployKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGitDeployKeyBadRequest creates a CreateGitDeployKeyBadRequest with default headers values
func NewCreateGitDeployKeyBadRequest() *CreateGitDeployKeyBadRequest {
	return &CreateGitDeployKeyBadRequest{}
}

/*CreateGitDeployKeyBadRequest handles this case with default header values.

Bad Request
*/
type CreateGitDeployKeyBadRequest struct {
	Payload *models.Error
}

func (o *CreateGitDeployKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/git/deploy_key][%d] createGitDeployKeyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGitDeployKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGitDeployKeyNotFound creates a CreateGitDeployKeyNotFound with default headers values
func NewCreateGitDeployKeyNotFound() *CreateGitDeployKeyNotFound {
	return &CreateGitDeployKeyNotFound{}
}

/*CreateGitDeployKeyNotFound handles this case with default header values.

Not Found
*/
type CreateGitDeployKeyNotFound struct {
	Payload *models.Error
}

func (o *CreateGitDeployKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/git/deploy_key][%d] createGitDeployKeyNotFound  %+v", 404, o.Payload)
}

func (o *CreateGitDeployKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGitDeployKeyConflict creates a CreateGitDeployKeyConflict with default headers values
func NewCreateGitDeployKeyConflict() *CreateGitDeployKeyConflict {
	return &CreateGitDeployKeyConflict{}
}

/*CreateGitDeployKeyConflict handles this case with default header values.

Resource Already Exists
*/
type CreateGitDeployKeyConflict struct {
	Payload *models.Error
}

func (o *CreateGitDeployKeyConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/git/deploy_key][%d] createGitDeployKeyConflict  %+v", 409, o.Payload)
}

func (o *CreateGitDeployKeyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
