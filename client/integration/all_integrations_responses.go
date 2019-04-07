// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// AllIntegrationsReader is a Reader for the AllIntegrations structure.
type AllIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllIntegrationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllIntegrationsOK creates a AllIntegrationsOK with default headers values
func NewAllIntegrationsOK() *AllIntegrationsOK {
	return &AllIntegrationsOK{}
}

/*AllIntegrationsOK handles this case with default header values.

Integration
*/
type AllIntegrationsOK struct {
	Payload []*models.Integration
}

func (o *AllIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /integrations][%d] allIntegrationsOK  %+v", 200, o.Payload)
}

func (o *AllIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllIntegrationsBadRequest creates a AllIntegrationsBadRequest with default headers values
func NewAllIntegrationsBadRequest() *AllIntegrationsBadRequest {
	return &AllIntegrationsBadRequest{}
}

/*AllIntegrationsBadRequest handles this case with default header values.

Bad Request
*/
type AllIntegrationsBadRequest struct {
	Payload *models.Error
}

func (o *AllIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /integrations][%d] allIntegrationsBadRequest  %+v", 400, o.Payload)
}

func (o *AllIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllIntegrationsNotFound creates a AllIntegrationsNotFound with default headers values
func NewAllIntegrationsNotFound() *AllIntegrationsNotFound {
	return &AllIntegrationsNotFound{}
}

/*AllIntegrationsNotFound handles this case with default header values.

Not Found
*/
type AllIntegrationsNotFound struct {
	Payload *models.Error
}

func (o *AllIntegrationsNotFound) Error() string {
	return fmt.Sprintf("[GET /integrations][%d] allIntegrationsNotFound  %+v", 404, o.Payload)
}

func (o *AllIntegrationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
