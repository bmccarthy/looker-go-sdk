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

// CreateScheduledPlanReader is a Reader for the CreateScheduledPlan structure.
type CreateScheduledPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScheduledPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateScheduledPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateScheduledPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateScheduledPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateScheduledPlanConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateScheduledPlanUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateScheduledPlanOK creates a CreateScheduledPlanOK with default headers values
func NewCreateScheduledPlanOK() *CreateScheduledPlanOK {
	return &CreateScheduledPlanOK{}
}

/*CreateScheduledPlanOK handles this case with default header values.

Scheduled Plan
*/
type CreateScheduledPlanOK struct {
	Payload *models.ScheduledPlan
}

func (o *CreateScheduledPlanOK) Error() string {
	return fmt.Sprintf("[POST /scheduled_plans][%d] createScheduledPlanOK  %+v", 200, o.Payload)
}

func (o *CreateScheduledPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledPlanBadRequest creates a CreateScheduledPlanBadRequest with default headers values
func NewCreateScheduledPlanBadRequest() *CreateScheduledPlanBadRequest {
	return &CreateScheduledPlanBadRequest{}
}

/*CreateScheduledPlanBadRequest handles this case with default header values.

Bad Request
*/
type CreateScheduledPlanBadRequest struct {
	Payload *models.Error
}

func (o *CreateScheduledPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /scheduled_plans][%d] createScheduledPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CreateScheduledPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledPlanNotFound creates a CreateScheduledPlanNotFound with default headers values
func NewCreateScheduledPlanNotFound() *CreateScheduledPlanNotFound {
	return &CreateScheduledPlanNotFound{}
}

/*CreateScheduledPlanNotFound handles this case with default header values.

Not Found
*/
type CreateScheduledPlanNotFound struct {
	Payload *models.Error
}

func (o *CreateScheduledPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /scheduled_plans][%d] createScheduledPlanNotFound  %+v", 404, o.Payload)
}

func (o *CreateScheduledPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledPlanConflict creates a CreateScheduledPlanConflict with default headers values
func NewCreateScheduledPlanConflict() *CreateScheduledPlanConflict {
	return &CreateScheduledPlanConflict{}
}

/*CreateScheduledPlanConflict handles this case with default header values.

Resource Already Exists
*/
type CreateScheduledPlanConflict struct {
	Payload *models.Error
}

func (o *CreateScheduledPlanConflict) Error() string {
	return fmt.Sprintf("[POST /scheduled_plans][%d] createScheduledPlanConflict  %+v", 409, o.Payload)
}

func (o *CreateScheduledPlanConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledPlanUnprocessableEntity creates a CreateScheduledPlanUnprocessableEntity with default headers values
func NewCreateScheduledPlanUnprocessableEntity() *CreateScheduledPlanUnprocessableEntity {
	return &CreateScheduledPlanUnprocessableEntity{}
}

/*CreateScheduledPlanUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateScheduledPlanUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateScheduledPlanUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /scheduled_plans][%d] createScheduledPlanUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateScheduledPlanUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
