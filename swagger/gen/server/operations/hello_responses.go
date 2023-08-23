// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ExcitingFrog/xuanwu/swagger/gen/models"
)

// HelloOKCode is the HTTP code returned for type HelloOK
const HelloOKCode int = 200

/*
HelloOK Hello

swagger:response helloOK
*/
type HelloOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewHelloOK creates HelloOK with default headers values
func NewHelloOK() *HelloOK {

	return &HelloOK{}
}

// WithPayload adds the payload to the hello o k response
func (o *HelloOK) WithPayload(payload string) *HelloOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello o k response
func (o *HelloOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// HelloBadRequestCode is the HTTP code returned for type HelloBadRequest
const HelloBadRequestCode int = 400

/*
HelloBadRequest BadRequest

swagger:response helloBadRequest
*/
type HelloBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewHelloBadRequest creates HelloBadRequest with default headers values
func NewHelloBadRequest() *HelloBadRequest {

	return &HelloBadRequest{}
}

// WithPayload adds the payload to the hello bad request response
func (o *HelloBadRequest) WithPayload(payload *models.ErrorResponse) *HelloBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello bad request response
func (o *HelloBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
