// Code generated by go-swagger; DO NOT EDIT.

package subjects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSubjectHandlerFunc turns a function with the right signature into a delete subject handler
type DeleteSubjectHandlerFunc func(DeleteSubjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSubjectHandlerFunc) Handle(params DeleteSubjectParams) middleware.Responder {
	return fn(params)
}

// DeleteSubjectHandler interface for that can handle valid delete subject params
type DeleteSubjectHandler interface {
	Handle(DeleteSubjectParams) middleware.Responder
}

// NewDeleteSubject creates a new http.Handler for the delete subject operation
func NewDeleteSubject(ctx *middleware.Context, handler DeleteSubjectHandler) *DeleteSubject {
	return &DeleteSubject{Context: ctx, Handler: handler}
}

/*
	DeleteSubject swagger:route DELETE /subjects/{id} subjects deleteSubject

# Delete a Subject

Deletes a subject from the database.
*/
type DeleteSubject struct {
	Context *middleware.Context
	Handler DeleteSubjectHandler
}

func (o *DeleteSubject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSubjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
