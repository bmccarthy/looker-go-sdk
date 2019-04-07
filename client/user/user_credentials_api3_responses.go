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

// UserCredentialsApi3Reader is a Reader for the UserCredentialsApi3 structure.
type UserCredentialsApi3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsApi3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCredentialsApi3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserCredentialsApi3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserCredentialsApi3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCredentialsApi3OK creates a UserCredentialsApi3OK with default headers values
func NewUserCredentialsApi3OK() *UserCredentialsApi3OK {
	return &UserCredentialsApi3OK{}
}

/*UserCredentialsApi3OK handles this case with default header values.

API 3 Credential
*/
type UserCredentialsApi3OK struct {
	Payload *models.CredentialsApi3
}

func (o *UserCredentialsApi3OK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] userCredentialsApi3OK  %+v", 200, o.Payload)
}

func (o *UserCredentialsApi3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsApi3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsApi3BadRequest creates a UserCredentialsApi3BadRequest with default headers values
func NewUserCredentialsApi3BadRequest() *UserCredentialsApi3BadRequest {
	return &UserCredentialsApi3BadRequest{}
}

/*UserCredentialsApi3BadRequest handles this case with default header values.

Bad Request
*/
type UserCredentialsApi3BadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsApi3BadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] userCredentialsApi3BadRequest  %+v", 400, o.Payload)
}

func (o *UserCredentialsApi3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsApi3NotFound creates a UserCredentialsApi3NotFound with default headers values
func NewUserCredentialsApi3NotFound() *UserCredentialsApi3NotFound {
	return &UserCredentialsApi3NotFound{}
}

/*UserCredentialsApi3NotFound handles this case with default header values.

Not Found
*/
type UserCredentialsApi3NotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsApi3NotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] userCredentialsApi3NotFound  %+v", 404, o.Payload)
}

func (o *UserCredentialsApi3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
