// Code generated by go-swagger; DO NOT EDIT.

package api_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

Access token with metadata.
*/
type LoginOK struct {
	Payload *models.AccessToken
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /login][%d] loginOK  %+v", 200, o.Payload)
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginBadRequest creates a LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {
	return &LoginBadRequest{}
}

/*LoginBadRequest handles this case with default header values.

Bad Request
*/
type LoginBadRequest struct {
	Payload *models.Error
}

func (o *LoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /login][%d] loginBadRequest  %+v", 400, o.Payload)
}

func (o *LoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginNotFound creates a LoginNotFound with default headers values
func NewLoginNotFound() *LoginNotFound {
	return &LoginNotFound{}
}

/*LoginNotFound handles this case with default header values.

Not Found
*/
type LoginNotFound struct {
	Payload *models.Error
}

func (o *LoginNotFound) Error() string {
	return fmt.Sprintf("[POST /login][%d] loginNotFound  %+v", 404, o.Payload)
}

func (o *LoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
