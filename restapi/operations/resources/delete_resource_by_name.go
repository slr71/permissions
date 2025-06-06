// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteResourceByNameHandlerFunc turns a function with the right signature into a delete resource by name handler
type DeleteResourceByNameHandlerFunc func(DeleteResourceByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteResourceByNameHandlerFunc) Handle(params DeleteResourceByNameParams) middleware.Responder {
	return fn(params)
}

// DeleteResourceByNameHandler interface for that can handle valid delete resource by name params
type DeleteResourceByNameHandler interface {
	Handle(DeleteResourceByNameParams) middleware.Responder
}

// NewDeleteResourceByName creates a new http.Handler for the delete resource by name operation
func NewDeleteResourceByName(ctx *middleware.Context, handler DeleteResourceByNameHandler) *DeleteResourceByName {
	return &DeleteResourceByName{Context: ctx, Handler: handler}
}

/*
	DeleteResourceByName swagger:route DELETE /resources resources deleteResourceByName

# Delete Resources by Name

Removes a resource from the database. A resource is a single item to which permissions may be applied. For example. The Discovery Environment app, Word Count, would be defined as a resource in the permissions service.
*/
type DeleteResourceByName struct {
	Context *middleware.Context
	Handler DeleteResourceByNameHandler
}

func (o *DeleteResourceByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteResourceByNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
