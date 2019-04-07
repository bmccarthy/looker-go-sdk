// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// UserCredentialsEmbedReader is a Reader for the UserCredentialsEmbed structure.
type UserCredentialsEmbedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsEmbedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCredentialsEmbedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserCredentialsEmbedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserCredentialsEmbedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCredentialsEmbedOK creates a UserCredentialsEmbedOK with default headers values
func NewUserCredentialsEmbedOK() *UserCredentialsEmbedOK {
	return &UserCredentialsEmbedOK{}
}

/*UserCredentialsEmbedOK handles this case with default header values.

Embedding Credential
*/
type UserCredentialsEmbedOK struct {
	Payload *models.CredentialsEmbed
}

func (o *UserCredentialsEmbedOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] userCredentialsEmbedOK  %+v", 200, o.Payload)
}

func (o *UserCredentialsEmbedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsEmbed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsEmbedBadRequest creates a UserCredentialsEmbedBadRequest with default headers values
func NewUserCredentialsEmbedBadRequest() *UserCredentialsEmbedBadRequest {
	return &UserCredentialsEmbedBadRequest{}
}

/*UserCredentialsEmbedBadRequest handles this case with default header values.

Bad Request
*/
type UserCredentialsEmbedBadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsEmbedBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] userCredentialsEmbedBadRequest  %+v", 400, o.Payload)
}

func (o *UserCredentialsEmbedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsEmbedNotFound creates a UserCredentialsEmbedNotFound with default headers values
func NewUserCredentialsEmbedNotFound() *UserCredentialsEmbedNotFound {
	return &UserCredentialsEmbedNotFound{}
}

/*UserCredentialsEmbedNotFound handles this case with default header values.

Not Found
*/
type UserCredentialsEmbedNotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsEmbedNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_embed/{credentials_embed_id}][%d] userCredentialsEmbedNotFound  %+v", 404, o.Payload)
}

func (o *UserCredentialsEmbedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
