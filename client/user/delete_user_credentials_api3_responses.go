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

// DeleteUserCredentialsApi3Reader is a Reader for the DeleteUserCredentialsApi3 structure.
type DeleteUserCredentialsApi3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialsApi3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserCredentialsApi3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserCredentialsApi3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserCredentialsApi3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserCredentialsApi3NoContent creates a DeleteUserCredentialsApi3NoContent with default headers values
func NewDeleteUserCredentialsApi3NoContent() *DeleteUserCredentialsApi3NoContent {
	return &DeleteUserCredentialsApi3NoContent{}
}

/*DeleteUserCredentialsApi3NoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsApi3NoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsApi3NoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] deleteUserCredentialsApi3NoContent  %+v", 204, o.Payload)
}

func (o *DeleteUserCredentialsApi3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsApi3BadRequest creates a DeleteUserCredentialsApi3BadRequest with default headers values
func NewDeleteUserCredentialsApi3BadRequest() *DeleteUserCredentialsApi3BadRequest {
	return &DeleteUserCredentialsApi3BadRequest{}
}

/*DeleteUserCredentialsApi3BadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserCredentialsApi3BadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsApi3BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] deleteUserCredentialsApi3BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserCredentialsApi3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsApi3NotFound creates a DeleteUserCredentialsApi3NotFound with default headers values
func NewDeleteUserCredentialsApi3NotFound() *DeleteUserCredentialsApi3NotFound {
	return &DeleteUserCredentialsApi3NotFound{}
}

/*DeleteUserCredentialsApi3NotFound handles this case with default header values.

Not Found
*/
type DeleteUserCredentialsApi3NotFound struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsApi3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_api3/{credentials_api3_id}][%d] deleteUserCredentialsApi3NotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserCredentialsApi3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
