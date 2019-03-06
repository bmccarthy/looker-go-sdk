// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// AllUserAttributesReader is a Reader for the AllUserAttributes structure.
type AllUserAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllUserAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllUserAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllUserAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllUserAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllUserAttributesOK creates a AllUserAttributesOK with default headers values
func NewAllUserAttributesOK() *AllUserAttributesOK {
	return &AllUserAttributesOK{}
}

/*AllUserAttributesOK handles this case with default header values.

User Attribute
*/
type AllUserAttributesOK struct {
	Payload []*models.UserAttribute
}

func (o *AllUserAttributesOK) Error() string {
	return fmt.Sprintf("[GET /user_attributes][%d] allUserAttributesOK  %+v", 200, o.Payload)
}

func (o *AllUserAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserAttributesBadRequest creates a AllUserAttributesBadRequest with default headers values
func NewAllUserAttributesBadRequest() *AllUserAttributesBadRequest {
	return &AllUserAttributesBadRequest{}
}

/*AllUserAttributesBadRequest handles this case with default header values.

Bad Request
*/
type AllUserAttributesBadRequest struct {
	Payload *models.Error
}

func (o *AllUserAttributesBadRequest) Error() string {
	return fmt.Sprintf("[GET /user_attributes][%d] allUserAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *AllUserAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserAttributesNotFound creates a AllUserAttributesNotFound with default headers values
func NewAllUserAttributesNotFound() *AllUserAttributesNotFound {
	return &AllUserAttributesNotFound{}
}

/*AllUserAttributesNotFound handles this case with default header values.

Not Found
*/
type AllUserAttributesNotFound struct {
	Payload *models.Error
}

func (o *AllUserAttributesNotFound) Error() string {
	return fmt.Sprintf("[GET /user_attributes][%d] allUserAttributesNotFound  %+v", 404, o.Payload)
}

func (o *AllUserAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}