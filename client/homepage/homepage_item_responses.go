// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// HomepageItemReader is a Reader for the HomepageItem structure.
type HomepageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HomepageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHomepageItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewHomepageItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewHomepageItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHomepageItemOK creates a HomepageItemOK with default headers values
func NewHomepageItemOK() *HomepageItemOK {
	return &HomepageItemOK{}
}

/*HomepageItemOK handles this case with default header values.

Homepage Item
*/
type HomepageItemOK struct {
	Payload *models.HomepageItem
}

func (o *HomepageItemOK) Error() string {
	return fmt.Sprintf("[GET /homepage_items/{homepage_item_id}][%d] homepageItemOK  %+v", 200, o.Payload)
}

func (o *HomepageItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomepageItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageItemBadRequest creates a HomepageItemBadRequest with default headers values
func NewHomepageItemBadRequest() *HomepageItemBadRequest {
	return &HomepageItemBadRequest{}
}

/*HomepageItemBadRequest handles this case with default header values.

Bad Request
*/
type HomepageItemBadRequest struct {
	Payload *models.Error
}

func (o *HomepageItemBadRequest) Error() string {
	return fmt.Sprintf("[GET /homepage_items/{homepage_item_id}][%d] homepageItemBadRequest  %+v", 400, o.Payload)
}

func (o *HomepageItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageItemNotFound creates a HomepageItemNotFound with default headers values
func NewHomepageItemNotFound() *HomepageItemNotFound {
	return &HomepageItemNotFound{}
}

/*HomepageItemNotFound handles this case with default header values.

Not Found
*/
type HomepageItemNotFound struct {
	Payload *models.Error
}

func (o *HomepageItemNotFound) Error() string {
	return fmt.Sprintf("[GET /homepage_items/{homepage_item_id}][%d] homepageItemNotFound  %+v", 404, o.Payload)
}

func (o *HomepageItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}