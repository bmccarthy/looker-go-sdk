// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// RoleUsersReader is a Reader for the RoleUsers structure.
type RoleUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRoleUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRoleUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRoleUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRoleUsersOK creates a RoleUsersOK with default headers values
func NewRoleUsersOK() *RoleUsersOK {
	return &RoleUsersOK{}
}

/*RoleUsersOK handles this case with default header values.

Users with role.
*/
type RoleUsersOK struct {
	Payload []*models.User
}

func (o *RoleUsersOK) Error() string {
	return fmt.Sprintf("[GET /roles/{role_id}/users][%d] roleUsersOK  %+v", 200, o.Payload)
}

func (o *RoleUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleUsersBadRequest creates a RoleUsersBadRequest with default headers values
func NewRoleUsersBadRequest() *RoleUsersBadRequest {
	return &RoleUsersBadRequest{}
}

/*RoleUsersBadRequest handles this case with default header values.

Bad Request
*/
type RoleUsersBadRequest struct {
	Payload *models.Error
}

func (o *RoleUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles/{role_id}/users][%d] roleUsersBadRequest  %+v", 400, o.Payload)
}

func (o *RoleUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleUsersNotFound creates a RoleUsersNotFound with default headers values
func NewRoleUsersNotFound() *RoleUsersNotFound {
	return &RoleUsersNotFound{}
}

/*RoleUsersNotFound handles this case with default header values.

Not Found
*/
type RoleUsersNotFound struct {
	Payload *models.Error
}

func (o *RoleUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /roles/{role_id}/users][%d] roleUsersNotFound  %+v", 404, o.Payload)
}

func (o *RoleUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
