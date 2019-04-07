// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// DeleteLookmlModelReader is a Reader for the DeleteLookmlModel structure.
type DeleteLookmlModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLookmlModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLookmlModelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLookmlModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLookmlModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLookmlModelNoContent creates a DeleteLookmlModelNoContent with default headers values
func NewDeleteLookmlModelNoContent() *DeleteLookmlModelNoContent {
	return &DeleteLookmlModelNoContent{}
}

/*DeleteLookmlModelNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteLookmlModelNoContent struct {
	Payload string
}

func (o *DeleteLookmlModelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lookml_models/{lookml_model_name}][%d] deleteLookmlModelNoContent  %+v", 204, o.Payload)
}

func (o *DeleteLookmlModelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLookmlModelBadRequest creates a DeleteLookmlModelBadRequest with default headers values
func NewDeleteLookmlModelBadRequest() *DeleteLookmlModelBadRequest {
	return &DeleteLookmlModelBadRequest{}
}

/*DeleteLookmlModelBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLookmlModelBadRequest struct {
	Payload *models.Error
}

func (o *DeleteLookmlModelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /lookml_models/{lookml_model_name}][%d] deleteLookmlModelBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLookmlModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLookmlModelNotFound creates a DeleteLookmlModelNotFound with default headers values
func NewDeleteLookmlModelNotFound() *DeleteLookmlModelNotFound {
	return &DeleteLookmlModelNotFound{}
}

/*DeleteLookmlModelNotFound handles this case with default header values.

Not Found
*/
type DeleteLookmlModelNotFound struct {
	Payload *models.Error
}

func (o *DeleteLookmlModelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /lookml_models/{lookml_model_name}][%d] deleteLookmlModelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLookmlModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
