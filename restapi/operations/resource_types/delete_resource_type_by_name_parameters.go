// Code generated by go-swagger; DO NOT EDIT.

package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteResourceTypeByNameParams creates a new DeleteResourceTypeByNameParams object
//
// There are no default values defined in the spec.
func NewDeleteResourceTypeByNameParams() DeleteResourceTypeByNameParams {

	return DeleteResourceTypeByNameParams{}
}

// DeleteResourceTypeByNameParams contains all the bound params for the delete resource type by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteResourceTypeByName
type DeleteResourceTypeByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The resource type name to search for.
	  Required: true
	  In: query
	*/
	ResourceTypeName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteResourceTypeByNameParams() beforehand.
func (o *DeleteResourceTypeByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qResourceTypeName, qhkResourceTypeName, _ := qs.GetOK("resource_type_name")
	if err := o.bindResourceTypeName(qResourceTypeName, qhkResourceTypeName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindResourceTypeName binds and validates parameter ResourceTypeName from query.
func (o *DeleteResourceTypeByNameParams) bindResourceTypeName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("resource_type_name", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("resource_type_name", "query", raw); err != nil {
		return err
	}
	o.ResourceTypeName = raw

	return nil
}
