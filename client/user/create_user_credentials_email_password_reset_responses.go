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

// CreateUserCredentialsEmailPasswordResetReader is a Reader for the CreateUserCredentialsEmailPasswordReset structure.
type CreateUserCredentialsEmailPasswordResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsEmailPasswordResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserCredentialsEmailPasswordResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserCredentialsEmailPasswordResetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateUserCredentialsEmailPasswordResetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialsEmailPasswordResetOK creates a CreateUserCredentialsEmailPasswordResetOK with default headers values
func NewCreateUserCredentialsEmailPasswordResetOK() *CreateUserCredentialsEmailPasswordResetOK {
	return &CreateUserCredentialsEmailPasswordResetOK{}
}

/*CreateUserCredentialsEmailPasswordResetOK handles this case with default header values.

email/password credential
*/
type CreateUserCredentialsEmailPasswordResetOK struct {
	Payload *models.CredentialsEmail
}

func (o *CreateUserCredentialsEmailPasswordResetOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email/password_reset][%d] createUserCredentialsEmailPasswordResetOK  %+v", 200, o.Payload)
}

func (o *CreateUserCredentialsEmailPasswordResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsEmail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailPasswordResetBadRequest creates a CreateUserCredentialsEmailPasswordResetBadRequest with default headers values
func NewCreateUserCredentialsEmailPasswordResetBadRequest() *CreateUserCredentialsEmailPasswordResetBadRequest {
	return &CreateUserCredentialsEmailPasswordResetBadRequest{}
}

/*CreateUserCredentialsEmailPasswordResetBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserCredentialsEmailPasswordResetBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsEmailPasswordResetBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email/password_reset][%d] createUserCredentialsEmailPasswordResetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserCredentialsEmailPasswordResetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailPasswordResetNotFound creates a CreateUserCredentialsEmailPasswordResetNotFound with default headers values
func NewCreateUserCredentialsEmailPasswordResetNotFound() *CreateUserCredentialsEmailPasswordResetNotFound {
	return &CreateUserCredentialsEmailPasswordResetNotFound{}
}

/*CreateUserCredentialsEmailPasswordResetNotFound handles this case with default header values.

Not Found
*/
type CreateUserCredentialsEmailPasswordResetNotFound struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsEmailPasswordResetNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email/password_reset][%d] createUserCredentialsEmailPasswordResetNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserCredentialsEmailPasswordResetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
