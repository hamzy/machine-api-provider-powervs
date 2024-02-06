// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_p_p_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudSppplacementgroupsPostParams creates a new PcloudSppplacementgroupsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudSppplacementgroupsPostParams() *PcloudSppplacementgroupsPostParams {
	return &PcloudSppplacementgroupsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudSppplacementgroupsPostParamsWithTimeout creates a new PcloudSppplacementgroupsPostParams object
// with the ability to set a timeout on a request.
func NewPcloudSppplacementgroupsPostParamsWithTimeout(timeout time.Duration) *PcloudSppplacementgroupsPostParams {
	return &PcloudSppplacementgroupsPostParams{
		timeout: timeout,
	}
}

// NewPcloudSppplacementgroupsPostParamsWithContext creates a new PcloudSppplacementgroupsPostParams object
// with the ability to set a context for a request.
func NewPcloudSppplacementgroupsPostParamsWithContext(ctx context.Context) *PcloudSppplacementgroupsPostParams {
	return &PcloudSppplacementgroupsPostParams{
		Context: ctx,
	}
}

// NewPcloudSppplacementgroupsPostParamsWithHTTPClient creates a new PcloudSppplacementgroupsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudSppplacementgroupsPostParamsWithHTTPClient(client *http.Client) *PcloudSppplacementgroupsPostParams {
	return &PcloudSppplacementgroupsPostParams{
		HTTPClient: client,
	}
}

/*
PcloudSppplacementgroupsPostParams contains all the parameters to send to the API endpoint

	for the pcloud sppplacementgroups post operation.

	Typically these are written to a http.Request.
*/
type PcloudSppplacementgroupsPostParams struct {

	/* Body.

	   Parameters for the creation of a Shared Processor Pool Placement Group
	*/
	Body *models.SPPPlacementGroupCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud sppplacementgroups post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudSppplacementgroupsPostParams) WithDefaults() *PcloudSppplacementgroupsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud sppplacementgroups post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudSppplacementgroupsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) WithTimeout(timeout time.Duration) *PcloudSppplacementgroupsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) WithContext(ctx context.Context) *PcloudSppplacementgroupsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) WithHTTPClient(client *http.Client) *PcloudSppplacementgroupsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) WithBody(body *models.SPPPlacementGroupCreate) *PcloudSppplacementgroupsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) SetBody(body *models.SPPPlacementGroupCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudSppplacementgroupsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud sppplacementgroups post params
func (o *PcloudSppplacementgroupsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudSppplacementgroupsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}