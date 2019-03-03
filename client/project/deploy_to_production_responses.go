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

// DeployToProductionReader is a Reader for the DeployToProduction structure.
type DeployToProductionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeployToProductionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeployToProductionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeployToProductionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeployToProductionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeployToProductionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeployToProductionOK creates a DeployToProductionOK with default headers values
func NewDeployToProductionOK() *DeployToProductionOK {
	return &DeployToProductionOK{}
}

/*DeployToProductionOK handles this case with default header values.

Project
*/
type DeployToProductionOK struct {
	Payload string
}

func (o *DeployToProductionOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/deploy_to_production][%d] deployToProductionOK  %+v", 200, o.Payload)
}

func (o *DeployToProductionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployToProductionNoContent creates a DeployToProductionNoContent with default headers values
func NewDeployToProductionNoContent() *DeployToProductionNoContent {
	return &DeployToProductionNoContent{}
}

/*DeployToProductionNoContent handles this case with default header values.

Returns 204 if project was successfully deployed to production, otherwise 400 with an error message
*/
type DeployToProductionNoContent struct {
}

func (o *DeployToProductionNoContent) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/deploy_to_production][%d] deployToProductionNoContent ", 204)
}

func (o *DeployToProductionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeployToProductionBadRequest creates a DeployToProductionBadRequest with default headers values
func NewDeployToProductionBadRequest() *DeployToProductionBadRequest {
	return &DeployToProductionBadRequest{}
}

/*DeployToProductionBadRequest handles this case with default header values.

Bad Request
*/
type DeployToProductionBadRequest struct {
	Payload *models.Error
}

func (o *DeployToProductionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/deploy_to_production][%d] deployToProductionBadRequest  %+v", 400, o.Payload)
}

func (o *DeployToProductionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployToProductionNotFound creates a DeployToProductionNotFound with default headers values
func NewDeployToProductionNotFound() *DeployToProductionNotFound {
	return &DeployToProductionNotFound{}
}

/*DeployToProductionNotFound handles this case with default header values.

Not Found
*/
type DeployToProductionNotFound struct {
	Payload *models.Error
}

func (o *DeployToProductionNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/deploy_to_production][%d] deployToProductionNotFound  %+v", 404, o.Payload)
}

func (o *DeployToProductionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
