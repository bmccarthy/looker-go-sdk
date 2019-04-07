// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// FetchAndParseSamlIdpMetadataReader is a Reader for the FetchAndParseSamlIdpMetadata structure.
type FetchAndParseSamlIdpMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchAndParseSamlIdpMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFetchAndParseSamlIdpMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFetchAndParseSamlIdpMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFetchAndParseSamlIdpMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFetchAndParseSamlIdpMetadataOK creates a FetchAndParseSamlIdpMetadataOK with default headers values
func NewFetchAndParseSamlIdpMetadataOK() *FetchAndParseSamlIdpMetadataOK {
	return &FetchAndParseSamlIdpMetadataOK{}
}

/*FetchAndParseSamlIdpMetadataOK handles this case with default header values.

Parse result
*/
type FetchAndParseSamlIdpMetadataOK struct {
	Payload *models.SamlMetadataParseResult
}

func (o *FetchAndParseSamlIdpMetadataOK) Error() string {
	return fmt.Sprintf("[POST /fetch_and_parse_saml_idp_metadata][%d] fetchAndParseSamlIdpMetadataOK  %+v", 200, o.Payload)
}

func (o *FetchAndParseSamlIdpMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SamlMetadataParseResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchAndParseSamlIdpMetadataBadRequest creates a FetchAndParseSamlIdpMetadataBadRequest with default headers values
func NewFetchAndParseSamlIdpMetadataBadRequest() *FetchAndParseSamlIdpMetadataBadRequest {
	return &FetchAndParseSamlIdpMetadataBadRequest{}
}

/*FetchAndParseSamlIdpMetadataBadRequest handles this case with default header values.

Bad Request
*/
type FetchAndParseSamlIdpMetadataBadRequest struct {
	Payload *models.Error
}

func (o *FetchAndParseSamlIdpMetadataBadRequest) Error() string {
	return fmt.Sprintf("[POST /fetch_and_parse_saml_idp_metadata][%d] fetchAndParseSamlIdpMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *FetchAndParseSamlIdpMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchAndParseSamlIdpMetadataNotFound creates a FetchAndParseSamlIdpMetadataNotFound with default headers values
func NewFetchAndParseSamlIdpMetadataNotFound() *FetchAndParseSamlIdpMetadataNotFound {
	return &FetchAndParseSamlIdpMetadataNotFound{}
}

/*FetchAndParseSamlIdpMetadataNotFound handles this case with default header values.

Not Found
*/
type FetchAndParseSamlIdpMetadataNotFound struct {
	Payload *models.Error
}

func (o *FetchAndParseSamlIdpMetadataNotFound) Error() string {
	return fmt.Sprintf("[POST /fetch_and_parse_saml_idp_metadata][%d] fetchAndParseSamlIdpMetadataNotFound  %+v", 404, o.Payload)
}

func (o *FetchAndParseSamlIdpMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
