// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VPNConnectionUpdate VPN Connection object to send during the update
// swagger:model VPNConnectionUpdate
type VPNConnectionUpdate struct {

	// unique identifier of IKEPolicy selected for this VPNConnection
	IkePolicy string `json:"ikePolicy,omitempty"`

	// unique identifier of IPSecPolicy selected for this VPNConnection
	IPSecPolicy string `json:"ipSecPolicy,omitempty"`

	// VPN Connection name
	Name string `json:"name,omitempty"`

	// peer gateway address
	// Format: ipv4
	PeerGatewayAddress PeerGatewayAddress `json:"peerGatewayAddress,omitempty"`

	// v p n connection update additional properties
	VPNConnectionUpdateAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VPNConnectionUpdate) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// unique identifier of IKEPolicy selected for this VPNConnection
		IkePolicy string `json:"ikePolicy,omitempty"`

		// unique identifier of IPSecPolicy selected for this VPNConnection
		IPSecPolicy string `json:"ipSecPolicy,omitempty"`

		// VPN Connection name
		Name string `json:"name,omitempty"`

		// peer gateway address
		// Format: ipv4
		PeerGatewayAddress PeerGatewayAddress `json:"peerGatewayAddress,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VPNConnectionUpdate

	rcv.IkePolicy = stage1.IkePolicy

	rcv.IPSecPolicy = stage1.IPSecPolicy

	rcv.Name = stage1.Name

	rcv.PeerGatewayAddress = stage1.PeerGatewayAddress

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ikePolicy")

	delete(stage2, "ipSecPolicy")

	delete(stage2, "name")

	delete(stage2, "peerGatewayAddress")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.VPNConnectionUpdateAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VPNConnectionUpdate) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// unique identifier of IKEPolicy selected for this VPNConnection
		IkePolicy string `json:"ikePolicy,omitempty"`

		// unique identifier of IPSecPolicy selected for this VPNConnection
		IPSecPolicy string `json:"ipSecPolicy,omitempty"`

		// VPN Connection name
		Name string `json:"name,omitempty"`

		// peer gateway address
		// Format: ipv4
		PeerGatewayAddress PeerGatewayAddress `json:"peerGatewayAddress,omitempty"`
	}

	stage1.IkePolicy = m.IkePolicy

	stage1.IPSecPolicy = m.IPSecPolicy

	stage1.Name = m.Name

	stage1.PeerGatewayAddress = m.PeerGatewayAddress

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VPNConnectionUpdateAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VPNConnectionUpdateAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this v p n connection update
func (m *VPNConnectionUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeerGatewayAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VPNConnectionUpdate) validatePeerGatewayAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PeerGatewayAddress) { // not required
		return nil
	}

	if err := m.PeerGatewayAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("peerGatewayAddress")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VPNConnectionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNConnectionUpdate) UnmarshalBinary(b []byte) error {
	var res VPNConnectionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
