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

// AcceptIntegrationHubLegalAgreementReader is a Reader for the AcceptIntegrationHubLegalAgreement structure.
type AcceptIntegrationHubLegalAgreementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptIntegrationHubLegalAgreementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAcceptIntegrationHubLegalAgreementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAcceptIntegrationHubLegalAgreementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAcceptIntegrationHubLegalAgreementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAcceptIntegrationHubLegalAgreementUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAcceptIntegrationHubLegalAgreementOK creates a AcceptIntegrationHubLegalAgreementOK with default headers values
func NewAcceptIntegrationHubLegalAgreementOK() *AcceptIntegrationHubLegalAgreementOK {
	return &AcceptIntegrationHubLegalAgreementOK{}
}

/*AcceptIntegrationHubLegalAgreementOK handles this case with default header values.

Integration hub
*/
type AcceptIntegrationHubLegalAgreementOK struct {
	Payload *models.IntegrationHub
}

func (o *AcceptIntegrationHubLegalAgreementOK) Error() string {
	return fmt.Sprintf("[POST /integration_hubs/{integration_hub_id}/accept_legal_agreement][%d] acceptIntegrationHubLegalAgreementOK  %+v", 200, o.Payload)
}

func (o *AcceptIntegrationHubLegalAgreementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationHub)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptIntegrationHubLegalAgreementBadRequest creates a AcceptIntegrationHubLegalAgreementBadRequest with default headers values
func NewAcceptIntegrationHubLegalAgreementBadRequest() *AcceptIntegrationHubLegalAgreementBadRequest {
	return &AcceptIntegrationHubLegalAgreementBadRequest{}
}

/*AcceptIntegrationHubLegalAgreementBadRequest handles this case with default header values.

Bad Request
*/
type AcceptIntegrationHubLegalAgreementBadRequest struct {
	Payload *models.Error
}

func (o *AcceptIntegrationHubLegalAgreementBadRequest) Error() string {
	return fmt.Sprintf("[POST /integration_hubs/{integration_hub_id}/accept_legal_agreement][%d] acceptIntegrationHubLegalAgreementBadRequest  %+v", 400, o.Payload)
}

func (o *AcceptIntegrationHubLegalAgreementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptIntegrationHubLegalAgreementNotFound creates a AcceptIntegrationHubLegalAgreementNotFound with default headers values
func NewAcceptIntegrationHubLegalAgreementNotFound() *AcceptIntegrationHubLegalAgreementNotFound {
	return &AcceptIntegrationHubLegalAgreementNotFound{}
}

/*AcceptIntegrationHubLegalAgreementNotFound handles this case with default header values.

Not Found
*/
type AcceptIntegrationHubLegalAgreementNotFound struct {
	Payload *models.Error
}

func (o *AcceptIntegrationHubLegalAgreementNotFound) Error() string {
	return fmt.Sprintf("[POST /integration_hubs/{integration_hub_id}/accept_legal_agreement][%d] acceptIntegrationHubLegalAgreementNotFound  %+v", 404, o.Payload)
}

func (o *AcceptIntegrationHubLegalAgreementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptIntegrationHubLegalAgreementUnprocessableEntity creates a AcceptIntegrationHubLegalAgreementUnprocessableEntity with default headers values
func NewAcceptIntegrationHubLegalAgreementUnprocessableEntity() *AcceptIntegrationHubLegalAgreementUnprocessableEntity {
	return &AcceptIntegrationHubLegalAgreementUnprocessableEntity{}
}

/*AcceptIntegrationHubLegalAgreementUnprocessableEntity handles this case with default header values.

Validation Error
*/
type AcceptIntegrationHubLegalAgreementUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *AcceptIntegrationHubLegalAgreementUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /integration_hubs/{integration_hub_id}/accept_legal_agreement][%d] acceptIntegrationHubLegalAgreementUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AcceptIntegrationHubLegalAgreementUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
