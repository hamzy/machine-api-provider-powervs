// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIkepoliciesPostReader is a Reader for the PcloudIkepoliciesPost structure.
type PcloudIkepoliciesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIkepoliciesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudIkepoliciesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudIkepoliciesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudIkepoliciesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudIkepoliciesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudIkepoliciesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudIkepoliciesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudIkepoliciesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudIkepoliciesPostOK creates a PcloudIkepoliciesPostOK with default headers values
func NewPcloudIkepoliciesPostOK() *PcloudIkepoliciesPostOK {
	return &PcloudIkepoliciesPostOK{}
}

/*PcloudIkepoliciesPostOK handles this case with default header values.

OK
*/
type PcloudIkepoliciesPostOK struct {
	Payload *models.IKEPolicy
}

func (o *PcloudIkepoliciesPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostOK  %+v", 200, o.Payload)
}

func (o *PcloudIkepoliciesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IKEPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostBadRequest creates a PcloudIkepoliciesPostBadRequest with default headers values
func NewPcloudIkepoliciesPostBadRequest() *PcloudIkepoliciesPostBadRequest {
	return &PcloudIkepoliciesPostBadRequest{}
}

/*PcloudIkepoliciesPostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudIkepoliciesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudIkepoliciesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostUnauthorized creates a PcloudIkepoliciesPostUnauthorized with default headers values
func NewPcloudIkepoliciesPostUnauthorized() *PcloudIkepoliciesPostUnauthorized {
	return &PcloudIkepoliciesPostUnauthorized{}
}

/*PcloudIkepoliciesPostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudIkepoliciesPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudIkepoliciesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostForbidden creates a PcloudIkepoliciesPostForbidden with default headers values
func NewPcloudIkepoliciesPostForbidden() *PcloudIkepoliciesPostForbidden {
	return &PcloudIkepoliciesPostForbidden{}
}

/*PcloudIkepoliciesPostForbidden handles this case with default header values.

Forbidden
*/
type PcloudIkepoliciesPostForbidden struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudIkepoliciesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostConflict creates a PcloudIkepoliciesPostConflict with default headers values
func NewPcloudIkepoliciesPostConflict() *PcloudIkepoliciesPostConflict {
	return &PcloudIkepoliciesPostConflict{}
}

/*PcloudIkepoliciesPostConflict handles this case with default header values.

Conflict
*/
type PcloudIkepoliciesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudIkepoliciesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostUnprocessableEntity creates a PcloudIkepoliciesPostUnprocessableEntity with default headers values
func NewPcloudIkepoliciesPostUnprocessableEntity() *PcloudIkepoliciesPostUnprocessableEntity {
	return &PcloudIkepoliciesPostUnprocessableEntity{}
}

/*PcloudIkepoliciesPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudIkepoliciesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudIkepoliciesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPostInternalServerError creates a PcloudIkepoliciesPostInternalServerError with default headers values
func NewPcloudIkepoliciesPostInternalServerError() *PcloudIkepoliciesPostInternalServerError {
	return &PcloudIkepoliciesPostInternalServerError{}
}

/*PcloudIkepoliciesPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudIkepoliciesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies][%d] pcloudIkepoliciesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudIkepoliciesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
