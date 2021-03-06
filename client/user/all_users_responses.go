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

// AllUsersReader is a Reader for the AllUsers structure.
type AllUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllUsersOK creates a AllUsersOK with default headers values
func NewAllUsersOK() *AllUsersOK {
	return &AllUsersOK{}
}

/*AllUsersOK handles this case with default header values.

All users.
*/
type AllUsersOK struct {
	Payload []*models.User
}

func (o *AllUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] allUsersOK  %+v", 200, o.Payload)
}

func (o *AllUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUsersBadRequest creates a AllUsersBadRequest with default headers values
func NewAllUsersBadRequest() *AllUsersBadRequest {
	return &AllUsersBadRequest{}
}

/*AllUsersBadRequest handles this case with default header values.

Bad Request
*/
type AllUsersBadRequest struct {
	Payload *models.Error
}

func (o *AllUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users][%d] allUsersBadRequest  %+v", 400, o.Payload)
}

func (o *AllUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUsersNotFound creates a AllUsersNotFound with default headers values
func NewAllUsersNotFound() *AllUsersNotFound {
	return &AllUsersNotFound{}
}

/*AllUsersNotFound handles this case with default header values.

Not Found
*/
type AllUsersNotFound struct {
	Payload *models.Error
}

func (o *AllUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /users][%d] allUsersNotFound  %+v", 404, o.Payload)
}

func (o *AllUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
