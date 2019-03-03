// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "looker-go-sdk/models"
)

// ContentFavoriteReader is a Reader for the ContentFavorite structure.
type ContentFavoriteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContentFavoriteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContentFavoriteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContentFavoriteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContentFavoriteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContentFavoriteOK creates a ContentFavoriteOK with default headers values
func NewContentFavoriteOK() *ContentFavoriteOK {
	return &ContentFavoriteOK{}
}

/*ContentFavoriteOK handles this case with default header values.

Favorite Content
*/
type ContentFavoriteOK struct {
	Payload *models.ContentFavorite
}

func (o *ContentFavoriteOK) Error() string {
	return fmt.Sprintf("[GET /content_favorite/{content_favorite_id}][%d] contentFavoriteOK  %+v", 200, o.Payload)
}

func (o *ContentFavoriteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentFavorite)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentFavoriteBadRequest creates a ContentFavoriteBadRequest with default headers values
func NewContentFavoriteBadRequest() *ContentFavoriteBadRequest {
	return &ContentFavoriteBadRequest{}
}

/*ContentFavoriteBadRequest handles this case with default header values.

Bad Request
*/
type ContentFavoriteBadRequest struct {
	Payload *models.Error
}

func (o *ContentFavoriteBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_favorite/{content_favorite_id}][%d] contentFavoriteBadRequest  %+v", 400, o.Payload)
}

func (o *ContentFavoriteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContentFavoriteNotFound creates a ContentFavoriteNotFound with default headers values
func NewContentFavoriteNotFound() *ContentFavoriteNotFound {
	return &ContentFavoriteNotFound{}
}

/*ContentFavoriteNotFound handles this case with default header values.

Not Found
*/
type ContentFavoriteNotFound struct {
	Payload *models.Error
}

func (o *ContentFavoriteNotFound) Error() string {
	return fmt.Sprintf("[GET /content_favorite/{content_favorite_id}][%d] contentFavoriteNotFound  %+v", 404, o.Payload)
}

func (o *ContentFavoriteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
