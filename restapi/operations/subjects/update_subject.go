// Code generated by go-swagger; DO NOT EDIT.

package subjects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateSubjectHandlerFunc turns a function with the right signature into a update subject handler
type UpdateSubjectHandlerFunc func(UpdateSubjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSubjectHandlerFunc) Handle(params UpdateSubjectParams) middleware.Responder {
	return fn(params)
}

// UpdateSubjectHandler interface for that can handle valid update subject params
type UpdateSubjectHandler interface {
	Handle(UpdateSubjectParams) middleware.Responder
}

// NewUpdateSubject creates a new http.Handler for the update subject operation
func NewUpdateSubject(ctx *middleware.Context, handler UpdateSubjectHandler) *UpdateSubject {
	return &UpdateSubject{Context: ctx, Handler: handler}
}

/*
	UpdateSubject swagger:route PUT /subjects/{id} subjects updateSubject

# Update a Subject

Updates a subject in the database. For full use of the permissions service, the subject should be present in Grouper and have the same subject ID in Grouper and the permissions service.
*/
type UpdateSubject struct {
	Context *middleware.Context
	Handler UpdateSubjectHandler
}

func (o *UpdateSubject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateSubjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
