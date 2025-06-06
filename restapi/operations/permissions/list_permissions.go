// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListPermissionsHandlerFunc turns a function with the right signature into a list permissions handler
type ListPermissionsHandlerFunc func(ListPermissionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPermissionsHandlerFunc) Handle(params ListPermissionsParams) middleware.Responder {
	return fn(params)
}

// ListPermissionsHandler interface for that can handle valid list permissions params
type ListPermissionsHandler interface {
	Handle(ListPermissionsParams) middleware.Responder
}

// NewListPermissions creates a new http.Handler for the list permissions operation
func NewListPermissions(ctx *middleware.Context, handler ListPermissionsHandler) *ListPermissions {
	return &ListPermissions{Context: ctx, Handler: handler}
}

/*
	ListPermissions swagger:route GET /permissions permissions listPermissions

# List Permissions

Lists all permissions in the permission database. The total number of permissions for all resources is likely to be quite large, so callers should be prepared to handle the response body. If this endpoint is used more frequently than anticipated, limit and offset parameters will be added for paging later.
*/
type ListPermissions struct {
	Context *middleware.Context
	Handler ListPermissionsHandler
}

func (o *ListPermissions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListPermissionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
