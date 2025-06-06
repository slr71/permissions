// Code generated by go-swagger; DO NOT EDIT.

package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteResourceTypesIDHandlerFunc turns a function with the right signature into a delete resource types ID handler
type DeleteResourceTypesIDHandlerFunc func(DeleteResourceTypesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteResourceTypesIDHandlerFunc) Handle(params DeleteResourceTypesIDParams) middleware.Responder {
	return fn(params)
}

// DeleteResourceTypesIDHandler interface for that can handle valid delete resource types ID params
type DeleteResourceTypesIDHandler interface {
	Handle(DeleteResourceTypesIDParams) middleware.Responder
}

// NewDeleteResourceTypesID creates a new http.Handler for the delete resource types ID operation
func NewDeleteResourceTypesID(ctx *middleware.Context, handler DeleteResourceTypesIDHandler) *DeleteResourceTypesID {
	return &DeleteResourceTypesID{Context: ctx, Handler: handler}
}

/*
	DeleteResourceTypesID swagger:route DELETE /resource_types/{id} resource_types deleteResourceTypesId

# Delete a Resource Type

Removes a resource type from the database. A resource type may only be removed if there are no resources associated with it.
*/
type DeleteResourceTypesID struct {
	Context *middleware.Context
	Handler DeleteResourceTypesIDHandler
}

func (o *DeleteResourceTypesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteResourceTypesIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
