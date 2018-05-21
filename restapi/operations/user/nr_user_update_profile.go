// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"Passport/models"
	"Passport/utils"
	"fmt"
	"time"
	"strconv"
	"bytes"
)

// NrUserUpdateProfileHandlerFunc turns a function with the right signature into a user update profile handler
type NrUserUpdateProfileHandlerFunc func(NrUserUpdateProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserUpdateProfileHandlerFunc) Handle(params NrUserUpdateProfileParams) middleware.Responder {
	return fn(params)
}

// NrUserUpdateProfileHandler interface for that can handle valid user update profile params
type NrUserUpdateProfileHandler interface {
	Handle(NrUserUpdateProfileParams) middleware.Responder
}

// NewNrUserUpdateProfile creates a new http.Handler for the user update profile operation
func NewNrUserUpdateProfile(ctx *middleware.Context, handler NrUserUpdateProfileHandler) *NrUserUpdateProfile {
	return &NrUserUpdateProfile{Context: ctx, Handler: handler}
}

/*NrUserUpdateProfile swagger:route POST /user/updateProfile User userUpdateProfile

修改用户资料接口

修改用户资料接口

*/
type NrUserUpdateProfile struct {
	Context *middleware.Context
	Handler NrUserUpdateProfileHandler
}

func (o *NrUserUpdateProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserUpdateProfileParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//res := o.Handler.Handle(Params) // actually handle the request

	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var res models.UserState
	var state models.State
	var data models.UserInfo

	db.Table(utils.T_USER).Where("id=?", utils.DecodeUserID(*Params.Body.Euid)).Where("status=0").Find(&data)

	// 不存在的用户
	if data.ID == 0 {
		state.UnmarshalBinary([]byte(utils.Response200(402, "用户不存在")))
		res.State = &state
		o.Context.Respond(rw, r, route.Produces, route, res)
		return
	}

	var sqlBuf bytes.Buffer
	sql.WriteString("UPDATE btk_User SET ")
	if len(Params.Body.BirthDay) > 0 {
		sqlBuf.WriteString("birth_day="+Params.Body.BirthDay+",")
	}
	if len(Params.Body.Blood) > 0 {
		sqlBuf.WriteString("blood="+Params.Body.Blood+",")
	}
	if len(Params.Body.Degree) > 0 {
		sqlBuf.WriteString("degree="+Params.Body.Degree+",")
	}
	if Params.Body.Gender > -1 {
		sqlBuf.WriteString("gender="+strconv.FormatInt(Params.Body.Gender, 10)+",")
	}
	if len(Params.Body.HomeArea) > 0 {
		sqlBuf.WriteString("home_area="+Params.Body.HomeArea+",")
	}
	if len(Params.Body.Interest) > 0 {
		sqlBuf.WriteString("interest="+Params.Body.Interest+",")
	}
	if len(Params.Body.Marriage) > 0 {
		sqlBuf.WriteString("marriage="+Params.Body.Marriage+",")
	}
	if len(Params.Body.NickName) > 0 {
		sqlBuf.WriteString("nick_name="+Params.Body.NickName+",")
	}
	if len(Params.Body.NowArea) > 0 {
		sqlBuf.WriteString("now_area="+Params.Body.NowArea+",")
	}
	if len(Params.Body.Profession) > 0 {
		sqlBuf.WriteString("profession="+Params.Body.Profession+",")
	}
	if len(Params.Body.Resume) > 0 {
		sqlBuf.WriteString("resume="+Params.Body.Resume+",")
	}
	if len(Params.Body.Salary) > 0 {
		sqlBuf.WriteString("salary="+Params.Body.Salary+",")
	}
	if len(Params.Body.School) > 0 {
		sqlBuf.WriteString("school="+Params.Body.School+",")
	}
	if len(Params.Body.Shape) > 0 {
		sqlBuf.WriteString("shape="+Params.Body.Shape+",")
	}
	if Params.Body.Stature > 0 {
		sqlBuf.WriteString("stature="+strconv.FormatInt(Params.Body.Stature, 10)+"cm"+",")
	}
	if Params.Body.Weight > 0 {
		sqlBuf.WriteString("weight="+strconv.FormatInt(Params.Body.Weight, 10)+"cm"+",")
	}
	sqlBuf.WriteString("update_at="+strconv.FormatInt(time.Now().Unix(), 10))
	sqlBuf.WriteString(" WHERE status=0 AND id=?")
	db.Raw(sqlBuf.String(), data.ID)

	state.UnmarshalBinary([]byte(utils.Response200(200, "修改成功")))
	res.State = &state

	data.Euid = *Params.Body.Euid
	data.ID = 0
	data.Avatar = utils.CompleteImage(data.Avatar)
	res.Data = &data

	o.Context.Respond(rw, r, route.Produces, route, res)

}
