// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// DeleteHomepageItemReader is a Reader for the DeleteHomepageItem structure.
type DeleteHomepageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHomepageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteHomepageItemNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteHomepageItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteHomepageItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteHomepageItemNoContent creates a DeleteHomepageItemNoContent with default headers values
func NewDeleteHomepageItemNoContent() *DeleteHomepageItemNoContent {
	return &DeleteHomepageItemNoContent{}
}

/*DeleteHomepageItemNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteHomepageItemNoContent struct {
	Payload string
}

func (o *DeleteHomepageItemNoContent) Error() string {
	return fmt.Sprintf("[DELETE /homepage_items/{homepage_item_id}][%d] deleteHomepageItemNoContent  %+v", 204, o.Payload)
}

func (o *DeleteHomepageItemNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHomepageItemBadRequest creates a DeleteHomepageItemBadRequest with default headers values
func NewDeleteHomepageItemBadRequest() *DeleteHomepageItemBadRequest {
	return &DeleteHomepageItemBadRequest{}
}

/*DeleteHomepageItemBadRequest handles this case with default header values.

Bad Request
*/
type DeleteHomepageItemBadRequest struct {
	Payload *models.Error
}

func (o *DeleteHomepageItemBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /homepage_items/{homepage_item_id}][%d] deleteHomepageItemBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteHomepageItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHomepageItemNotFound creates a DeleteHomepageItemNotFound with default headers values
func NewDeleteHomepageItemNotFound() *DeleteHomepageItemNotFound {
	return &DeleteHomepageItemNotFound{}
}

/*DeleteHomepageItemNotFound handles this case with default header values.

Not Found
*/
type DeleteHomepageItemNotFound struct {
	Payload *models.Error
}

func (o *DeleteHomepageItemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /homepage_items/{homepage_item_id}][%d] deleteHomepageItemNotFound  %+v", 404, o.Payload)
}

func (o *DeleteHomepageItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
