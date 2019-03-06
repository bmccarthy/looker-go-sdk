// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// TestLdapConfigConnectionReader is a Reader for the TestLdapConfigConnection structure.
type TestLdapConfigConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLdapConfigConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestLdapConfigConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTestLdapConfigConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTestLdapConfigConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewTestLdapConfigConnectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestLdapConfigConnectionOK creates a TestLdapConfigConnectionOK with default headers values
func NewTestLdapConfigConnectionOK() *TestLdapConfigConnectionOK {
	return &TestLdapConfigConnectionOK{}
}

/*TestLdapConfigConnectionOK handles this case with default header values.

Result info.
*/
type TestLdapConfigConnectionOK struct {
	Payload *models.LDAPConfigTestResult
}

func (o *TestLdapConfigConnectionOK) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_connection][%d] testLdapConfigConnectionOK  %+v", 200, o.Payload)
}

func (o *TestLdapConfigConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPConfigTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigConnectionBadRequest creates a TestLdapConfigConnectionBadRequest with default headers values
func NewTestLdapConfigConnectionBadRequest() *TestLdapConfigConnectionBadRequest {
	return &TestLdapConfigConnectionBadRequest{}
}

/*TestLdapConfigConnectionBadRequest handles this case with default header values.

Bad Request
*/
type TestLdapConfigConnectionBadRequest struct {
	Payload *models.Error
}

func (o *TestLdapConfigConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_connection][%d] testLdapConfigConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *TestLdapConfigConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigConnectionNotFound creates a TestLdapConfigConnectionNotFound with default headers values
func NewTestLdapConfigConnectionNotFound() *TestLdapConfigConnectionNotFound {
	return &TestLdapConfigConnectionNotFound{}
}

/*TestLdapConfigConnectionNotFound handles this case with default header values.

Not Found
*/
type TestLdapConfigConnectionNotFound struct {
	Payload *models.Error
}

func (o *TestLdapConfigConnectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_connection][%d] testLdapConfigConnectionNotFound  %+v", 404, o.Payload)
}

func (o *TestLdapConfigConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigConnectionUnprocessableEntity creates a TestLdapConfigConnectionUnprocessableEntity with default headers values
func NewTestLdapConfigConnectionUnprocessableEntity() *TestLdapConfigConnectionUnprocessableEntity {
	return &TestLdapConfigConnectionUnprocessableEntity{}
}

/*TestLdapConfigConnectionUnprocessableEntity handles this case with default header values.

Validation Error
*/
type TestLdapConfigConnectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *TestLdapConfigConnectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_connection][%d] testLdapConfigConnectionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TestLdapConfigConnectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}