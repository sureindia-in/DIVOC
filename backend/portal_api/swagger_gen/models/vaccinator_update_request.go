// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VaccinatorUpdateRequest vaccinator update request
//
// swagger:model VaccinatorUpdateRequest
type VaccinatorUpdateRequest []*VaccinatorUpdateRequestItems0

// Validate validates this vaccinator update request
func (m VaccinatorUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// VaccinatorUpdateRequestItems0 vaccinator update request items0
//
// swagger:model VaccinatorUpdateRequestItems0
type VaccinatorUpdateRequestItems0 struct {

	// average rating
	AverageRating float64 `json:"averageRating,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// facility ids
	FacilityIds []string `json:"facilityIds"`

	// mobile number
	// Max Length: 10
	// Min Length: 10
	MobileNumber string `json:"mobileNumber,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// national identifier
	NationalIdentifier string `json:"nationalIdentifier,omitempty"`

	// osid
	Osid interface{} `json:"osid,omitempty"`

	// programs
	Programs []*VaccinatorUpdateRequestItems0ProgramsItems0 `json:"programs"`

	// signature string
	SignatureString string `json:"signatureString,omitempty"`

	// status
	// Enum: [Active Inactive]
	Status string `json:"status,omitempty"`

	// training certificate
	TrainingCertificate string `json:"trainingCertificate,omitempty"`
}

// Validate validates this vaccinator update request items0
func (m *VaccinatorUpdateRequestItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMobileNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrograms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaccinatorUpdateRequestItems0) validateMobileNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.MobileNumber) { // not required
		return nil
	}

	if err := validate.MinLength("mobileNumber", "body", string(m.MobileNumber), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("mobileNumber", "body", string(m.MobileNumber), 10); err != nil {
		return err
	}

	return nil
}

func (m *VaccinatorUpdateRequestItems0) validatePrograms(formats strfmt.Registry) error {

	if swag.IsZero(m.Programs) { // not required
		return nil
	}

	for i := 0; i < len(m.Programs); i++ {
		if swag.IsZero(m.Programs[i]) { // not required
			continue
		}

		if m.Programs[i] != nil {
			if err := m.Programs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("programs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var vaccinatorUpdateRequestItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vaccinatorUpdateRequestItems0TypeStatusPropEnum = append(vaccinatorUpdateRequestItems0TypeStatusPropEnum, v)
	}
}

const (

	// VaccinatorUpdateRequestItems0StatusActive captures enum value "Active"
	VaccinatorUpdateRequestItems0StatusActive string = "Active"

	// VaccinatorUpdateRequestItems0StatusInactive captures enum value "Inactive"
	VaccinatorUpdateRequestItems0StatusInactive string = "Inactive"
)

// prop value enum
func (m *VaccinatorUpdateRequestItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vaccinatorUpdateRequestItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VaccinatorUpdateRequestItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaccinatorUpdateRequestItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaccinatorUpdateRequestItems0) UnmarshalBinary(b []byte) error {
	var res VaccinatorUpdateRequestItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VaccinatorUpdateRequestItems0ProgramsItems0 vaccinator update request items0 programs items0
//
// swagger:model VaccinatorUpdateRequestItems0ProgramsItems0
type VaccinatorUpdateRequestItems0ProgramsItems0 struct {

	// certified
	Certified bool `json:"certified,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// status
	// Enum: [Active Inactive]
	Status string `json:"status,omitempty"`
}

// Validate validates this vaccinator update request items0 programs items0
func (m *VaccinatorUpdateRequestItems0ProgramsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vaccinatorUpdateRequestItems0ProgramsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vaccinatorUpdateRequestItems0ProgramsItems0TypeStatusPropEnum = append(vaccinatorUpdateRequestItems0ProgramsItems0TypeStatusPropEnum, v)
	}
}

const (

	// VaccinatorUpdateRequestItems0ProgramsItems0StatusActive captures enum value "Active"
	VaccinatorUpdateRequestItems0ProgramsItems0StatusActive string = "Active"

	// VaccinatorUpdateRequestItems0ProgramsItems0StatusInactive captures enum value "Inactive"
	VaccinatorUpdateRequestItems0ProgramsItems0StatusInactive string = "Inactive"
)

// prop value enum
func (m *VaccinatorUpdateRequestItems0ProgramsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vaccinatorUpdateRequestItems0ProgramsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VaccinatorUpdateRequestItems0ProgramsItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaccinatorUpdateRequestItems0ProgramsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaccinatorUpdateRequestItems0ProgramsItems0) UnmarshalBinary(b []byte) error {
	var res VaccinatorUpdateRequestItems0ProgramsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
