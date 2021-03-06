// Code generated by go-swagger; DO NOT EDIT.

package passport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"Passport/models"
)

// PassportLoginOKCode is the HTTP code returned for type PassportLoginOK
const PassportLoginOKCode int = 200

/*PassportLoginOK 登录成功，返回登录信息

swagger:response passportLoginOK
*/
type PassportLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginState `json:"body,omitempty"`
}

// NewPassportLoginOK creates PassportLoginOK with default headers values
func NewPassportLoginOK() *PassportLoginOK {
	return &PassportLoginOK{}
}

// WithPayload adds the payload to the passport login o k response
func (o *PassportLoginOK) WithPayload(payload *models.LoginState) *PassportLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the passport login o k response
func (o *PassportLoginOK) SetPayload(payload *models.LoginState) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PassportLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrPassportLoginDefault error

swagger:response passportLoginDefault
*/
type NrPassportLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrPassportLoginDefault creates NrPassportLoginDefault with default headers values
func NewNrPassportLoginDefault(code int) *NrPassportLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &NrPassportLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the passport login default response
func (o *NrPassportLoginDefault) WithStatusCode(code int) *NrPassportLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the passport login default response
func (o *NrPassportLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the passport login default response
func (o *NrPassportLoginDefault) WithPayload(payload *models.Error) *NrPassportLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the passport login default response
func (o *NrPassportLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrPassportLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
