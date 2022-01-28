// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesNetworksGetallReader is a Reader for the PcloudPvminstancesNetworksGetall structure.
type PcloudPvminstancesNetworksGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesNetworksGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesNetworksGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudPvminstancesNetworksGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesNetworksGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksGetallOK creates a PcloudPvminstancesNetworksGetallOK with default headers values
func NewPcloudPvminstancesNetworksGetallOK() *PcloudPvminstancesNetworksGetallOK {
	return &PcloudPvminstancesNetworksGetallOK{}
}

/*PcloudPvminstancesNetworksGetallOK handles this case with default header values.

OK
*/
type PcloudPvminstancesNetworksGetallOK struct {
	Payload *models.PVMInstanceNetworks
}

func (o *PcloudPvminstancesNetworksGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallBadRequest creates a PcloudPvminstancesNetworksGetallBadRequest with default headers values
func NewPcloudPvminstancesNetworksGetallBadRequest() *PcloudPvminstancesNetworksGetallBadRequest {
	return &PcloudPvminstancesNetworksGetallBadRequest{}
}

/*PcloudPvminstancesNetworksGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesNetworksGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallUnauthorized creates a PcloudPvminstancesNetworksGetallUnauthorized with default headers values
func NewPcloudPvminstancesNetworksGetallUnauthorized() *PcloudPvminstancesNetworksGetallUnauthorized {
	return &PcloudPvminstancesNetworksGetallUnauthorized{}
}

/*PcloudPvminstancesNetworksGetallUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudPvminstancesNetworksGetallUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetallInternalServerError creates a PcloudPvminstancesNetworksGetallInternalServerError with default headers values
func NewPcloudPvminstancesNetworksGetallInternalServerError() *PcloudPvminstancesNetworksGetallInternalServerError {
	return &PcloudPvminstancesNetworksGetallInternalServerError{}
}

/*PcloudPvminstancesNetworksGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesNetworksGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
