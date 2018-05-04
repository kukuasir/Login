// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserUpdatePwdParamsBody user update pwd params body
// swagger:model userUpdatePwdParamsBody
type UserUpdatePwdParamsBody struct {

	// 客户端类型
	Client string `json:"client,omitempty"`

	// 用户ID
	// Required: true
	Euid *string `json:"euid"`

	// 唯一识别号
	Imei string `json:"imei,omitempty"`

	// 新密码
	// Required: true
	NewPwd *string `json:"new_pwd"`

	// 旧密码
	// Required: true
	OldPwd *string `json:"old_pwd"`

	// 手机号
	Phone string `json:"phone,omitempty"`

	// 时间戳(加密串要用到,供服务端验证，简单防刷)
	Ts int64 `json:"ts,omitempty"`

	// 加密串
	Val string `json:"val,omitempty"`

	// 版本号
	Version string `json:"version,omitempty"`
}

// Validate validates this user update pwd params body
func (m *UserUpdatePwdParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewPwd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOldPwd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUpdatePwdParamsBody) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *UserUpdatePwdParamsBody) validateNewPwd(formats strfmt.Registry) error {

	if err := validate.Required("new_pwd", "body", m.NewPwd); err != nil {
		return err
	}

	return nil
}

func (m *UserUpdatePwdParamsBody) validateOldPwd(formats strfmt.Registry) error {

	if err := validate.Required("old_pwd", "body", m.OldPwd); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdatePwdParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdatePwdParamsBody) UnmarshalBinary(b []byte) error {
	var res UserUpdatePwdParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
