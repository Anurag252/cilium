// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EndpointState State of endpoint
//
// swagger:model EndpointState
type EndpointState string

const (

	// EndpointStateWaitingForIdentity captures enum value "waiting-for-identity"
	EndpointStateWaitingForIdentity EndpointState = "waiting-for-identity"

	// EndpointStateNotReady captures enum value "not-ready"
	EndpointStateNotReady EndpointState = "not-ready"

	// EndpointStateWaitingToRegenerate captures enum value "waiting-to-regenerate"
	EndpointStateWaitingToRegenerate EndpointState = "waiting-to-regenerate"

	// EndpointStateRegenerating captures enum value "regenerating"
	EndpointStateRegenerating EndpointState = "regenerating"

	// EndpointStateRestoring captures enum value "restoring"
	EndpointStateRestoring EndpointState = "restoring"

	// EndpointStateReady captures enum value "ready"
	EndpointStateReady EndpointState = "ready"

	// EndpointStateDisconnecting captures enum value "disconnecting"
	EndpointStateDisconnecting EndpointState = "disconnecting"

	// EndpointStateDisconnected captures enum value "disconnected"
	EndpointStateDisconnected EndpointState = "disconnected"

	// EndpointStateInvalid captures enum value "invalid"
	EndpointStateInvalid EndpointState = "invalid"
)

// for schema
var endpointStateEnum []interface{}

func init() {
	var res []EndpointState
	if err := json.Unmarshal([]byte(`["waiting-for-identity","not-ready","waiting-to-regenerate","regenerating","restoring","ready","disconnecting","disconnected","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointStateEnum = append(endpointStateEnum, v)
	}
}

func (m EndpointState) validateEndpointStateEnum(path, location string, value EndpointState) error {
	if err := validate.EnumCase(path, location, value, endpointStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this endpoint state
func (m EndpointState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEndpointStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
