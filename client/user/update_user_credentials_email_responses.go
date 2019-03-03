// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// UpdateUserCredentialsEmailReader is a Reader for the UpdateUserCredentialsEmail structure.
type UpdateUserCredentialsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserCredentialsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserCredentialsEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateUserCredentialsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateUserCredentialsEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateUserCredentialsEmailUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserCredentialsEmailOK creates a UpdateUserCredentialsEmailOK with default headers values
func NewUpdateUserCredentialsEmailOK() *UpdateUserCredentialsEmailOK {
	return &UpdateUserCredentialsEmailOK{}
}

/*UpdateUserCredentialsEmailOK handles this case with default header values.

Email/Password Credential
*/
type UpdateUserCredentialsEmailOK struct {
	Payload *models.CredentialsEmail
}

func (o *UpdateUserCredentialsEmailOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}/credentials_email][%d] updateUserCredentialsEmailOK  %+v", 200, o.Payload)
}

func (o *UpdateUserCredentialsEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsEmail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserCredentialsEmailBadRequest creates a UpdateUserCredentialsEmailBadRequest with default headers values
func NewUpdateUserCredentialsEmailBadRequest() *UpdateUserCredentialsEmailBadRequest {
	return &UpdateUserCredentialsEmailBadRequest{}
}

/*UpdateUserCredentialsEmailBadRequest handles this case with default header values.

Bad Request
*/
type UpdateUserCredentialsEmailBadRequest struct {
	Payload *models.Error
}

func (o *UpdateUserCredentialsEmailBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}/credentials_email][%d] updateUserCredentialsEmailBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserCredentialsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserCredentialsEmailNotFound creates a UpdateUserCredentialsEmailNotFound with default headers values
func NewUpdateUserCredentialsEmailNotFound() *UpdateUserCredentialsEmailNotFound {
	return &UpdateUserCredentialsEmailNotFound{}
}

/*UpdateUserCredentialsEmailNotFound handles this case with default header values.

Not Found
*/
type UpdateUserCredentialsEmailNotFound struct {
	Payload *models.Error
}

func (o *UpdateUserCredentialsEmailNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}/credentials_email][%d] updateUserCredentialsEmailNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserCredentialsEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserCredentialsEmailUnprocessableEntity creates a UpdateUserCredentialsEmailUnprocessableEntity with default headers values
func NewUpdateUserCredentialsEmailUnprocessableEntity() *UpdateUserCredentialsEmailUnprocessableEntity {
	return &UpdateUserCredentialsEmailUnprocessableEntity{}
}

/*UpdateUserCredentialsEmailUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateUserCredentialsEmailUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateUserCredentialsEmailUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}/credentials_email][%d] updateUserCredentialsEmailUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserCredentialsEmailUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
