// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SettingsFlow SettingsFlow SettingsFlow SettingsFlow Flow represents a Settings Flow
//
// This flow is used when an identity wants to update settings
// (e.g. profile data, passwords, ...) in a selfservice manner.
//
// We recommend reading the [User Settings Documentation](../self-service/flows/user-settings)
//
// swagger:model settingsFlow
type SettingsFlow struct {

	// Active, if set, contains the registration method that is being used. It is initially
	// not set.
	Active string `json:"active,omitempty"`

	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to update the setting,
	// a new flow has to be initiated.
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expires_at"`

	// id
	// Required: true
	// Format: uuid4
	ID *UUID `json:"id"`

	// identity
	// Required: true
	Identity *Identity `json:"identity"`

	// IssuedAt is the time (UTC) when the flow occurred.
	// Required: true
	// Format: date-time
	IssuedAt *strfmt.DateTime `json:"issued_at"`

	// messages
	Messages Messages `json:"messages,omitempty"`

	// Methods contains context for all enabled registration methods. If a settings flow has been
	// processed, but for example the first name is empty, this will contain error messages.
	// Required: true
	Methods map[string]SettingsFlowMethod `json:"methods"`

	// RequestURL is the initial URL that was requested from ORY Kratos. It can be used
	// to forward information contained in the URL's path or query for example.
	// Required: true
	RequestURL *string `json:"request_url"`

	// state
	// Required: true
	State *State `json:"state"`

	// type
	Type Type `json:"type,omitempty"`
}

// Validate validates this settings flow
func (m *SettingsFlow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsFlow) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SettingsFlow) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) validateIdentity(formats strfmt.Registry) error {

	if err := validate.Required("identity", "body", m.Identity); err != nil {
		return err
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) validateIssuedAt(formats strfmt.Registry) error {

	if err := validate.Required("issued_at", "body", m.IssuedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SettingsFlow) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	if err := m.Messages.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("messages")
		}
		return err
	}

	return nil
}

func (m *SettingsFlow) validateMethods(formats strfmt.Registry) error {

	if err := validate.Required("methods", "body", m.Methods); err != nil {
		return err
	}

	for k := range m.Methods {

		if err := validate.Required("methods"+"."+k, "body", m.Methods[k]); err != nil {
			return err
		}
		if val, ok := m.Methods[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SettingsFlow) validateRequestURL(formats strfmt.Registry) error {

	if err := validate.Required("request_url", "body", m.RequestURL); err != nil {
		return err
	}

	return nil
}

func (m *SettingsFlow) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this settings flow based on the context it is used
func (m *SettingsFlow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsFlow) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {
		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.Identity != nil {
		if err := m.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Messages.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("messages")
		}
		return err
	}

	return nil
}

func (m *SettingsFlow) contextValidateMethods(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("methods", "body", m.Methods); err != nil {
		return err
	}

	for k := range m.Methods {

		if val, ok := m.Methods[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SettingsFlow) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *SettingsFlow) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsFlow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsFlow) UnmarshalBinary(b []byte) error {
	var res SettingsFlow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
