// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// IPAMMetadata Additional IPAM metadata when allocating IPs
//
//
// swagger:model IPAMMetadata
type IPAMMetadata map[string]string

// Validate validates this IP a m metadata
func (m IPAMMetadata) Validate(formats strfmt.Registry) error {
	return nil
}
