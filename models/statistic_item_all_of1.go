// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatisticItemAllOf1 statistic item all of1
//
// swagger:model statisticItemAllOf1
type StatisticItemAllOf1 struct {

	// set to "StatisticItem"
	// Enum: [StatisticItem]
	Type interface{} `json:"type,omitempty"`

	// The value of item
	Value string `json:"value,omitempty"`
}

// Validate validates this statistic item all of1
func (m *StatisticItemAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatisticItemAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatisticItemAllOf1) UnmarshalBinary(b []byte) error {
	var res StatisticItemAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
