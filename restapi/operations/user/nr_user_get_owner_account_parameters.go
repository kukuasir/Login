// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrUserGetOwnerAccountParams creates a new NrUserGetOwnerAccountParams object
// with the default values initialized.
func NewNrUserGetOwnerAccountParams() NrUserGetOwnerAccountParams {
	var ()
	return NrUserGetOwnerAccountParams{}
}

// NrUserGetOwnerAccountParams contains all the bound params for the user get owner account operation
// typically these are obtained from a http.Request
//
// swagger:parameters /user/getOwnerAccount
type NrUserGetOwnerAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*加密后的用户ID
	  Required: true
	  In: query
	*/
	Euid string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrUserGetOwnerAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEuid, qhkEuid, _ := qs.GetOK("euid")
	if err := o.bindEuid(qEuid, qhkEuid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrUserGetOwnerAccountParams) bindEuid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("euid", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("euid", "query", raw); err != nil {
		return err
	}

	o.Euid = raw

	return nil
}
