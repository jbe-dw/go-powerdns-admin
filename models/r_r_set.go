// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RRSet RRSet
//
// This represents a Resource Record Set (all records with the same name and type).
//
// swagger:model RRSet
type RRSet struct {

	// MUST be added when updating the RRSet. Must be REPLACE or DELETE. With DELETE, all existing RRs matching name and type will be deleted, including all comments. With REPLACE: when records is present, all existing RRs matching name and type will be deleted, and then new records given in records will be created. If no records are left, any existing comments will be deleted as well. When comments is present, all existing comments for the RRs matching name and type will be deleted, and then new comments given in comments will be created.
	// Required: true
	Changetype *string `json:"changetype"`

	// List of Comment. Must be empty when changetype is set to DELETE. An empty list results in deletion of all comments. modified_at is optional and defaults to the current server time.
	Comments []*Comment `json:"comments"`

	// Name for record set (e.g. “www.powerdns.com.”)
	// Required: true
	Name *string `json:"name"`

	// All records in this RRSet. When updating Records, this is the list of new records (replacing the old ones). Must be empty when changetype is set to DELETE. An empty list results in deletion of all records (and comments).
	// Required: true
	Records []*Record `json:"records"`

	// DNS TTL of the records, in seconds. MUST NOT be included when changetype is set to “DELETE”.
	// Required: true
	TTL *int64 `json:"ttl"`

	// Type of this record (e.g. “A”, “PTR”, “MX”)
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this r r set
func (m *RRSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangetype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
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

func (m *RRSet) validateChangetype(formats strfmt.Registry) error {

	if err := validate.Required("changetype", "body", m.Changetype); err != nil {
		return err
	}

	return nil
}

func (m *RRSet) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RRSet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RRSet) validateRecords(formats strfmt.Registry) error {

	if err := validate.Required("records", "body", m.Records); err != nil {
		return err
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RRSet) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("ttl", "body", m.TTL); err != nil {
		return err
	}

	return nil
}

func (m *RRSet) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RRSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RRSet) UnmarshalBinary(b []byte) error {
	var res RRSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
