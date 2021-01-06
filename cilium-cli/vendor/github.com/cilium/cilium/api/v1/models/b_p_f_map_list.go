// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BPFMapList List of BPF Maps
//
// swagger:model BPFMapList
type BPFMapList struct {

	// Array of open BPF map lists
	Maps []*BPFMap `json:"maps"`
}

// Validate validates this b p f map list
func (m *BPFMapList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BPFMapList) validateMaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Maps) { // not required
		return nil
	}

	for i := 0; i < len(m.Maps); i++ {
		if swag.IsZero(m.Maps[i]) { // not required
			continue
		}

		if m.Maps[i] != nil {
			if err := m.Maps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BPFMapList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BPFMapList) UnmarshalBinary(b []byte) error {
	var res BPFMapList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
