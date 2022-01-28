// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudImagesGetallParams creates a new PcloudImagesGetallParams object
// with the default values initialized.
func NewPcloudImagesGetallParams() *PcloudImagesGetallParams {
	var ()
	return &PcloudImagesGetallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudImagesGetallParamsWithTimeout creates a new PcloudImagesGetallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudImagesGetallParamsWithTimeout(timeout time.Duration) *PcloudImagesGetallParams {
	var ()
	return &PcloudImagesGetallParams{

		timeout: timeout,
	}
}

// NewPcloudImagesGetallParamsWithContext creates a new PcloudImagesGetallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudImagesGetallParamsWithContext(ctx context.Context) *PcloudImagesGetallParams {
	var ()
	return &PcloudImagesGetallParams{

		Context: ctx,
	}
}

// NewPcloudImagesGetallParamsWithHTTPClient creates a new PcloudImagesGetallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudImagesGetallParamsWithHTTPClient(client *http.Client) *PcloudImagesGetallParams {
	var ()
	return &PcloudImagesGetallParams{
		HTTPClient: client,
	}
}

/*PcloudImagesGetallParams contains all the parameters to send to the API endpoint
for the pcloud images getall operation typically these are written to a http.Request
*/
type PcloudImagesGetallParams struct {

	/*Sap
	  Include SAP images with get available stock images

	*/
	Sap *bool
	/*Vtl
	  Include VTL images with get available stock images

	*/
	Vtl *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud images getall params
func (o *PcloudImagesGetallParams) WithTimeout(timeout time.Duration) *PcloudImagesGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud images getall params
func (o *PcloudImagesGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud images getall params
func (o *PcloudImagesGetallParams) WithContext(ctx context.Context) *PcloudImagesGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud images getall params
func (o *PcloudImagesGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud images getall params
func (o *PcloudImagesGetallParams) WithHTTPClient(client *http.Client) *PcloudImagesGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud images getall params
func (o *PcloudImagesGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSap adds the sap to the pcloud images getall params
func (o *PcloudImagesGetallParams) WithSap(sap *bool) *PcloudImagesGetallParams {
	o.SetSap(sap)
	return o
}

// SetSap adds the sap to the pcloud images getall params
func (o *PcloudImagesGetallParams) SetSap(sap *bool) {
	o.Sap = sap
}

// WithVtl adds the vtl to the pcloud images getall params
func (o *PcloudImagesGetallParams) WithVtl(vtl *bool) *PcloudImagesGetallParams {
	o.SetVtl(vtl)
	return o
}

// SetVtl adds the vtl to the pcloud images getall params
func (o *PcloudImagesGetallParams) SetVtl(vtl *bool) {
	o.Vtl = vtl
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudImagesGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Sap != nil {

		// query param sap
		var qrSap bool
		if o.Sap != nil {
			qrSap = *o.Sap
		}
		qSap := swag.FormatBool(qrSap)
		if qSap != "" {
			if err := r.SetQueryParam("sap", qSap); err != nil {
				return err
			}
		}

	}

	if o.Vtl != nil {

		// query param vtl
		var qrVtl bool
		if o.Vtl != nil {
			qrVtl = *o.Vtl
		}
		qVtl := swag.FormatBool(qrVtl)
		if qVtl != "" {
			if err := r.SetQueryParam("vtl", qVtl); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
