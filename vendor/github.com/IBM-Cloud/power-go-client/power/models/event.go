// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
// swagger:model Event
type Event struct {

	// Type of action for this event
	// Required: true
	Action *string `json:"action"`

	// ID of the Activity
	// Required: true
	EventID *string `json:"eventID"`

	// Level of the event (notice, info, warning, error)
	// Required: true
	// Enum: [notice info warning error]
	Level *string `json:"level"`

	// The (translated) message of the event
	// Required: true
	Message *string `json:"message"`

	// Any metadata associated with the event
	Metadata interface{} `json:"metadata,omitempty"`

	// Type of resource for this event
	// Required: true
	Resource *string `json:"resource"`

	// Time of activity in ISO 8601 - RFC3339
	// Required: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time"`

	// Time of activity in unix epoch
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// user
	User *EventUser `json:"user,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventID", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

var eventTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["notice","info","warning","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeLevelPropEnum = append(eventTypeLevelPropEnum, v)
	}
}

const (

	// EventLevelNotice captures enum value "notice"
	EventLevelNotice string = "notice"

	// EventLevelInfo captures enum value "info"
	EventLevelInfo string = "info"

	// EventLevelWarning captures enum value "warning"
	EventLevelWarning string = "warning"

	// EventLevelError captures enum value "error"
	EventLevelError string = "error"
)

// prop value enum
func (m *Event) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, eventTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Event) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", *m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
