// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// DeleteScheduledPlanReader is a Reader for the DeleteScheduledPlan structure.
type DeleteScheduledPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScheduledPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteScheduledPlanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteScheduledPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteScheduledPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteScheduledPlanNoContent creates a DeleteScheduledPlanNoContent with default headers values
func NewDeleteScheduledPlanNoContent() *DeleteScheduledPlanNoContent {
	return &DeleteScheduledPlanNoContent{}
}

/*DeleteScheduledPlanNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteScheduledPlanNoContent struct {
	Payload string
}

func (o *DeleteScheduledPlanNoContent) Error() string {
	return fmt.Sprintf("[DELETE /scheduled_plans/{scheduled_plan_id}][%d] deleteScheduledPlanNoContent  %+v", 204, o.Payload)
}

func (o *DeleteScheduledPlanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduledPlanBadRequest creates a DeleteScheduledPlanBadRequest with default headers values
func NewDeleteScheduledPlanBadRequest() *DeleteScheduledPlanBadRequest {
	return &DeleteScheduledPlanBadRequest{}
}

/*DeleteScheduledPlanBadRequest handles this case with default header values.

Bad Request
*/
type DeleteScheduledPlanBadRequest struct {
	Payload *models.Error
}

func (o *DeleteScheduledPlanBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /scheduled_plans/{scheduled_plan_id}][%d] deleteScheduledPlanBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScheduledPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduledPlanNotFound creates a DeleteScheduledPlanNotFound with default headers values
func NewDeleteScheduledPlanNotFound() *DeleteScheduledPlanNotFound {
	return &DeleteScheduledPlanNotFound{}
}

/*DeleteScheduledPlanNotFound handles this case with default header values.

Not Found
*/
type DeleteScheduledPlanNotFound struct {
	Payload *models.Error
}

func (o *DeleteScheduledPlanNotFound) Error() string {
	return fmt.Sprintf("[DELETE /scheduled_plans/{scheduled_plan_id}][%d] deleteScheduledPlanNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduledPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
