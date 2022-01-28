// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesVolumesActionPostReader is a Reader for the PcloudCloudinstancesVolumesActionPost structure.
type PcloudCloudinstancesVolumesActionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesActionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPcloudCloudinstancesVolumesActionPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudinstancesVolumesActionPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudCloudinstancesVolumesActionPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudCloudinstancesVolumesActionPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudinstancesVolumesActionPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesActionPostAccepted creates a PcloudCloudinstancesVolumesActionPostAccepted with default headers values
func NewPcloudCloudinstancesVolumesActionPostAccepted() *PcloudCloudinstancesVolumesActionPostAccepted {
	return &PcloudCloudinstancesVolumesActionPostAccepted{}
}

/*PcloudCloudinstancesVolumesActionPostAccepted handles this case with default header values.

Accepted
*/
type PcloudCloudinstancesVolumesActionPostAccepted struct {
	Payload models.Object
}

func (o *PcloudCloudinstancesVolumesActionPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/action][%d] pcloudCloudinstancesVolumesActionPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudinstancesVolumesActionPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesActionPostBadRequest creates a PcloudCloudinstancesVolumesActionPostBadRequest with default headers values
func NewPcloudCloudinstancesVolumesActionPostBadRequest() *PcloudCloudinstancesVolumesActionPostBadRequest {
	return &PcloudCloudinstancesVolumesActionPostBadRequest{}
}

/*PcloudCloudinstancesVolumesActionPostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesActionPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesActionPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/action][%d] pcloudCloudinstancesVolumesActionPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesActionPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesActionPostUnauthorized creates a PcloudCloudinstancesVolumesActionPostUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesActionPostUnauthorized() *PcloudCloudinstancesVolumesActionPostUnauthorized {
	return &PcloudCloudinstancesVolumesActionPostUnauthorized{}
}

/*PcloudCloudinstancesVolumesActionPostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesActionPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesActionPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/action][%d] pcloudCloudinstancesVolumesActionPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesActionPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesActionPostNotFound creates a PcloudCloudinstancesVolumesActionPostNotFound with default headers values
func NewPcloudCloudinstancesVolumesActionPostNotFound() *PcloudCloudinstancesVolumesActionPostNotFound {
	return &PcloudCloudinstancesVolumesActionPostNotFound{}
}

/*PcloudCloudinstancesVolumesActionPostNotFound handles this case with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesActionPostNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesActionPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/action][%d] pcloudCloudinstancesVolumesActionPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesActionPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesActionPostInternalServerError creates a PcloudCloudinstancesVolumesActionPostInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesActionPostInternalServerError() *PcloudCloudinstancesVolumesActionPostInternalServerError {
	return &PcloudCloudinstancesVolumesActionPostInternalServerError{}
}

/*PcloudCloudinstancesVolumesActionPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesActionPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesActionPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}/action][%d] pcloudCloudinstancesVolumesActionPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesActionPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
