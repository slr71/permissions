// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cyverse-de/permissions/models"
)

// BySubjectOKCode is the HTTP code returned for type BySubjectOK
const BySubjectOKCode int = 200

/*
BySubjectOK OK

swagger:response bySubjectOK
*/
type BySubjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.PermissionList `json:"body,omitempty"`
}

// NewBySubjectOK creates BySubjectOK with default headers values
func NewBySubjectOK() *BySubjectOK {

	return &BySubjectOK{}
}

// WithPayload adds the payload to the by subject o k response
func (o *BySubjectOK) WithPayload(payload *models.PermissionList) *BySubjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject o k response
func (o *BySubjectOK) SetPayload(payload *models.PermissionList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BySubjectBadRequestCode is the HTTP code returned for type BySubjectBadRequest
const BySubjectBadRequestCode int = 400

/*
BySubjectBadRequest Bad Request

swagger:response bySubjectBadRequest
*/
type BySubjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewBySubjectBadRequest creates BySubjectBadRequest with default headers values
func NewBySubjectBadRequest() *BySubjectBadRequest {

	return &BySubjectBadRequest{}
}

// WithPayload adds the payload to the by subject bad request response
func (o *BySubjectBadRequest) WithPayload(payload *models.ErrorOut) *BySubjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject bad request response
func (o *BySubjectBadRequest) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BySubjectInternalServerErrorCode is the HTTP code returned for type BySubjectInternalServerError
const BySubjectInternalServerErrorCode int = 500

/*
BySubjectInternalServerError Internal Server Error

swagger:response bySubjectInternalServerError
*/
type BySubjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewBySubjectInternalServerError creates BySubjectInternalServerError with default headers values
func NewBySubjectInternalServerError() *BySubjectInternalServerError {

	return &BySubjectInternalServerError{}
}

// WithPayload adds the payload to the by subject internal server error response
func (o *BySubjectInternalServerError) WithPayload(payload *models.ErrorOut) *BySubjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the by subject internal server error response
func (o *BySubjectInternalServerError) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BySubjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
