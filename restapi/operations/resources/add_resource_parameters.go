// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/cyverse-de/permissions/models"
)

// NewAddResourceParams creates a new AddResourceParams object
//
// There are no default values defined in the spec.
func NewAddResourceParams() AddResourceParams {

	return AddResourceParams{}
}

// AddResourceParams contains all the bound params for the add resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters addResource
type AddResourceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The new resource to add.
	  Required: true
	  In: body
	*/
	ResourceIn *models.ResourceIn
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddResourceParams() beforehand.
func (o *AddResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ResourceIn
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("resourceIn", "body", ""))
			} else {
				res = append(res, errors.NewParseError("resourceIn", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ResourceIn = &body
			}
		}
	} else {
		res = append(res, errors.Required("resourceIn", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
