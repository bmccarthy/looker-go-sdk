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

// AllUserSessionsReader is a Reader for the AllUserSessions structure.
type AllUserSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllUserSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllUserSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllUserSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllUserSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllUserSessionsOK creates a AllUserSessionsOK with default headers values
func NewAllUserSessionsOK() *AllUserSessionsOK {
	return &AllUserSessionsOK{}
}

/*AllUserSessionsOK handles this case with default header values.

Web Login Session
*/
type AllUserSessionsOK struct {
	Payload []*models.Session
}

func (o *AllUserSessionsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions][%d] allUserSessionsOK  %+v", 200, o.Payload)
}

func (o *AllUserSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserSessionsBadRequest creates a AllUserSessionsBadRequest with default headers values
func NewAllUserSessionsBadRequest() *AllUserSessionsBadRequest {
	return &AllUserSessionsBadRequest{}
}

/*AllUserSessionsBadRequest handles this case with default header values.

Bad Request
*/
type AllUserSessionsBadRequest struct {
	Payload *models.Error
}

func (o *AllUserSessionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions][%d] allUserSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *AllUserSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserSessionsNotFound creates a AllUserSessionsNotFound with default headers values
func NewAllUserSessionsNotFound() *AllUserSessionsNotFound {
	return &AllUserSessionsNotFound{}
}

/*AllUserSessionsNotFound handles this case with default header values.

Not Found
*/
type AllUserSessionsNotFound struct {
	Payload *models.Error
}

func (o *AllUserSessionsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/sessions][%d] allUserSessionsNotFound  %+v", 404, o.Payload)
}

func (o *AllUserSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
