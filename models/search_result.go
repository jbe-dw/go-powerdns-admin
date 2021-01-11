// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchResult SearchResult
//
// swagger:model SearchResult
type SearchResult struct {

	// content
	Content string `json:"content,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// set to one of "record, zone, comment"
	ObjectType string `json:"object_type,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`

	// zone id
	ZoneID string `json:"zone_id,omitempty"`
}

// Validate validates this search result
func (m *SearchResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResult) UnmarshalBinary(b []byte) error {
	var res SearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
