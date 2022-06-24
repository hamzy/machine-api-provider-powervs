// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud volume groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud volume groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudVolumegroupsActionPost(params *PcloudVolumegroupsActionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsActionPostAccepted, error)

	PcloudVolumegroupsDelete(params *PcloudVolumegroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsDeleteAccepted, error)

	PcloudVolumegroupsGet(params *PcloudVolumegroupsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetOK, error)

	PcloudVolumegroupsGetDetails(params *PcloudVolumegroupsGetDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetDetailsOK, error)

	PcloudVolumegroupsGetall(params *PcloudVolumegroupsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetallOK, error)

	PcloudVolumegroupsGetallDetails(params *PcloudVolumegroupsGetallDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetallDetailsOK, error)

	PcloudVolumegroupsPost(params *PcloudVolumegroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsPostAccepted, error)

	PcloudVolumegroupsPut(params *PcloudVolumegroupsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsPutAccepted, error)

	PcloudVolumegroupsRemoteCopyRelationshipsGet(params *PcloudVolumegroupsRemoteCopyRelationshipsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsRemoteCopyRelationshipsGetOK, error)

	PcloudVolumegroupsStorageDetailsGet(params *PcloudVolumegroupsStorageDetailsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsStorageDetailsGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PcloudVolumegroupsActionPost performs an action start stop reset on a volume group
*/
func (a *Client) PcloudVolumegroupsActionPost(params *PcloudVolumegroupsActionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsActionPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsActionPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.action.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsActionPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsActionPostAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.action.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsDelete deletes a cloud instance volume group
*/
func (a *Client) PcloudVolumegroupsDelete(params *PcloudVolumegroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsDeleteAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsGet gets volume group
*/
func (a *Client) PcloudVolumegroupsGet(params *PcloudVolumegroupsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsGetDetails gets volume group details
*/
func (a *Client) PcloudVolumegroupsGetDetails(params *PcloudVolumegroupsGetDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsGetDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.getDetails",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsGetDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsGetDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.getDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsGetall gets all volume groups
*/
func (a *Client) PcloudVolumegroupsGetall(params *PcloudVolumegroupsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsGetallDetails gets all volume groups with details
*/
func (a *Client) PcloudVolumegroupsGetallDetails(params *PcloudVolumegroupsGetallDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsGetallDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsGetallDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.getallDetails",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsGetallDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsGetallDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.getallDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsPost creates a new volume group
*/
func (a *Client) PcloudVolumegroupsPost(params *PcloudVolumegroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsPostAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsPut updates the volume group
*/
func (a *Client) PcloudVolumegroupsPut(params *PcloudVolumegroupsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsPutAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsRemoteCopyRelationshipsGet gets remote copy relationships of the volume belonging to volume group
*/
func (a *Client) PcloudVolumegroupsRemoteCopyRelationshipsGet(params *PcloudVolumegroupsRemoteCopyRelationshipsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsRemoteCopyRelationshipsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsRemoteCopyRelationshipsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.remoteCopyRelationships.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/remote-copy-relationships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsRemoteCopyRelationshipsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsRemoteCopyRelationshipsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.remoteCopyRelationships.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudVolumegroupsStorageDetailsGet gets storage details of volume group
*/
func (a *Client) PcloudVolumegroupsStorageDetailsGet(params *PcloudVolumegroupsStorageDetailsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudVolumegroupsStorageDetailsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudVolumegroupsStorageDetailsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.volumegroups.storageDetails.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/storage-details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudVolumegroupsStorageDetailsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudVolumegroupsStorageDetailsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.volumegroups.storageDetails.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
