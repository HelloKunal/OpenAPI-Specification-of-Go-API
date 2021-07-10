// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"example.com/swagger/models"
)

// UpdateOneOKCode is the HTTP code returned for type UpdateOneOK
const UpdateOneOKCode int = 200

/*UpdateOneOK OK

swagger:response updateOneOK
*/
type UpdateOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewUpdateOneOK creates UpdateOneOK with default headers values
func NewUpdateOneOK() *UpdateOneOK {

	return &UpdateOneOK{}
}

// WithPayload adds the payload to the update one o k response
func (o *UpdateOneOK) WithPayload(payload *models.Item) *UpdateOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update one o k response
func (o *UpdateOneOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateOneDefault error

swagger:response updateOneDefault
*/
type UpdateOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateOneDefault creates UpdateOneDefault with default headers values
func NewUpdateOneDefault(code int) *UpdateOneDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update one default response
func (o *UpdateOneDefault) WithStatusCode(code int) *UpdateOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update one default response
func (o *UpdateOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update one default response
func (o *UpdateOneDefault) WithPayload(payload *models.Error) *UpdateOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update one default response
func (o *UpdateOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
