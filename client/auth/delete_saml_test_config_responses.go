// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// DeleteSamlTestConfigReader is a Reader for the DeleteSamlTestConfig structure.
type DeleteSamlTestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSamlTestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSamlTestConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSamlTestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSamlTestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSamlTestConfigNoContent creates a DeleteSamlTestConfigNoContent with default headers values
func NewDeleteSamlTestConfigNoContent() *DeleteSamlTestConfigNoContent {
	return &DeleteSamlTestConfigNoContent{}
}

/*DeleteSamlTestConfigNoContent handles this case with default header values.

Test config succssfully deleted.
*/
type DeleteSamlTestConfigNoContent struct {
	Payload string
}

func (o *DeleteSamlTestConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /saml_test_configs/{test_slug}][%d] deleteSamlTestConfigNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSamlTestConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSamlTestConfigBadRequest creates a DeleteSamlTestConfigBadRequest with default headers values
func NewDeleteSamlTestConfigBadRequest() *DeleteSamlTestConfigBadRequest {
	return &DeleteSamlTestConfigBadRequest{}
}

/*DeleteSamlTestConfigBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSamlTestConfigBadRequest struct {
	Payload *models.Error
}

func (o *DeleteSamlTestConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /saml_test_configs/{test_slug}][%d] deleteSamlTestConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSamlTestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSamlTestConfigNotFound creates a DeleteSamlTestConfigNotFound with default headers values
func NewDeleteSamlTestConfigNotFound() *DeleteSamlTestConfigNotFound {
	return &DeleteSamlTestConfigNotFound{}
}

/*DeleteSamlTestConfigNotFound handles this case with default header values.

Not Found
*/
type DeleteSamlTestConfigNotFound struct {
	Payload *models.Error
}

func (o *DeleteSamlTestConfigNotFound) Error() string {
	return fmt.Sprintf("[DELETE /saml_test_configs/{test_slug}][%d] deleteSamlTestConfigNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSamlTestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
