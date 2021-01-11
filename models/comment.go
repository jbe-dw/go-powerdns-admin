// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Comment Comment
//
// A comment about an RRSet.
//
// swagger:model Comment
type Comment struct {

	// Name of an account that added the comment
	Account string `json:"account,omitempty"`

	// The actual comment
	Content string `json:"content,omitempty"`

	// Timestamp of the last change to the comment
	ModifiedAt int64 `json:"modified_at,omitempty"`
}

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
