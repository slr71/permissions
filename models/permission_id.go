package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PermissionID The internal permission identifier.
// swagger:model permission_id
type PermissionID string

// Validate validates this permission id
func (m PermissionID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 36); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 36); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}