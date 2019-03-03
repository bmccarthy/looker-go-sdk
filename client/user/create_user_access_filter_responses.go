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

// CreateUserAccessFilterReader is a Reader for the CreateUserAccessFilter structure.
type CreateUserAccessFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserAccessFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserAccessFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserAccessFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateUserAccessFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateUserAccessFilterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateUserAccessFilterUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserAccessFilterOK creates a CreateUserAccessFilterOK with default headers values
func NewCreateUserAccessFilterOK() *CreateUserAccessFilterOK {
	return &CreateUserAccessFilterOK{}
}

/*CreateUserAccessFilterOK handles this case with default header values.

Access Filter
*/
type CreateUserAccessFilterOK struct {
	Payload *models.AccessFilter
}

func (o *CreateUserAccessFilterOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/access_filters][%d] createUserAccessFilterOK  %+v", 200, o.Payload)
}

func (o *CreateUserAccessFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAccessFilterBadRequest creates a CreateUserAccessFilterBadRequest with default headers values
func NewCreateUserAccessFilterBadRequest() *CreateUserAccessFilterBadRequest {
	return &CreateUserAccessFilterBadRequest{}
}

/*CreateUserAccessFilterBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserAccessFilterBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserAccessFilterBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/access_filters][%d] createUserAccessFilterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserAccessFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAccessFilterNotFound creates a CreateUserAccessFilterNotFound with default headers values
func NewCreateUserAccessFilterNotFound() *CreateUserAccessFilterNotFound {
	return &CreateUserAccessFilterNotFound{}
}

/*CreateUserAccessFilterNotFound handles this case with default header values.

Not Found
*/
type CreateUserAccessFilterNotFound struct {
	Payload *models.Error
}

func (o *CreateUserAccessFilterNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/access_filters][%d] createUserAccessFilterNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserAccessFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAccessFilterConflict creates a CreateUserAccessFilterConflict with default headers values
func NewCreateUserAccessFilterConflict() *CreateUserAccessFilterConflict {
	return &CreateUserAccessFilterConflict{}
}

/*CreateUserAccessFilterConflict handles this case with default header values.

Resource Already Exists
*/
type CreateUserAccessFilterConflict struct {
	Payload *models.Error
}

func (o *CreateUserAccessFilterConflict) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/access_filters][%d] createUserAccessFilterConflict  %+v", 409, o.Payload)
}

func (o *CreateUserAccessFilterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAccessFilterUnprocessableEntity creates a CreateUserAccessFilterUnprocessableEntity with default headers values
func NewCreateUserAccessFilterUnprocessableEntity() *CreateUserAccessFilterUnprocessableEntity {
	return &CreateUserAccessFilterUnprocessableEntity{}
}

/*CreateUserAccessFilterUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateUserAccessFilterUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserAccessFilterUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/access_filters][%d] createUserAccessFilterUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserAccessFilterUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
