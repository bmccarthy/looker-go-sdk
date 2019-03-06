// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateUserCredentialsAPIReader is a Reader for the CreateUserCredentialsAPI structure.
type CreateUserCredentialsAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserCredentialsAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserCredentialsAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateUserCredentialsAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialsAPIOK creates a CreateUserCredentialsAPIOK with default headers values
func NewCreateUserCredentialsAPIOK() *CreateUserCredentialsAPIOK {
	return &CreateUserCredentialsAPIOK{}
}

/*CreateUserCredentialsAPIOK handles this case with default header values.

API Credential
*/
type CreateUserCredentialsAPIOK struct {
	Payload *models.CredentialsAPI
}

func (o *CreateUserCredentialsAPIOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api][%d] createUserCredentialsApiOK  %+v", 200, o.Payload)
}

func (o *CreateUserCredentialsAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsAPI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsAPIBadRequest creates a CreateUserCredentialsAPIBadRequest with default headers values
func NewCreateUserCredentialsAPIBadRequest() *CreateUserCredentialsAPIBadRequest {
	return &CreateUserCredentialsAPIBadRequest{}
}

/*CreateUserCredentialsAPIBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserCredentialsAPIBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsAPIBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api][%d] createUserCredentialsApiBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserCredentialsAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsAPINotFound creates a CreateUserCredentialsAPINotFound with default headers values
func NewCreateUserCredentialsAPINotFound() *CreateUserCredentialsAPINotFound {
	return &CreateUserCredentialsAPINotFound{}
}

/*CreateUserCredentialsAPINotFound handles this case with default header values.

Not Found
*/
type CreateUserCredentialsAPINotFound struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsAPINotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api][%d] createUserCredentialsApiNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserCredentialsAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}