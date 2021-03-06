// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"net/url"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"Passport/models"
	"Passport/utils"
	"fmt"
	"time"
	"path"
)

// NrUserUpdateAvatarHandlerFunc turns a function with the right signature into a user update avatar handler
type NrUserUpdateAvatarHandlerFunc func(NrUserUpdateAvatarParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserUpdateAvatarHandlerFunc) Handle(params NrUserUpdateAvatarParams) middleware.Responder {
	return fn(params)
}

// NrUserUpdateAvatarHandler interface for that can handle valid user update avatar params
type NrUserUpdateAvatarHandler interface {
	Handle(NrUserUpdateAvatarParams) middleware.Responder
}

// NewNrUserUpdateAvatar creates a new http.Handler for the user update avatar operation
func NewNrUserUpdateAvatar(ctx *middleware.Context, handler NrUserUpdateAvatarHandler) *NrUserUpdateAvatar {
	return &NrUserUpdateAvatar{Context: ctx, Handler: handler}
}

/*NrUserUpdateAvatar swagger:route POST /user/updateAvatar User userUpdateAvatar

修改头像API

*/
type NrUserUpdateAvatar struct {
	Context *middleware.Context
	Handler NrUserUpdateAvatarHandler
}

func (o *NrUserUpdateAvatar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserUpdateAvatarParams()

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

	var res models.AvatarState
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

	var filename string
	if len(Params.Body.Avatar) > 0 {
		u, _ := url.Parse(Params.Body.Avatar)
		filename = path.Base(u.Path)
	}

	sql := "UPDATE btk_User SET avatar = ?, update_at = ? WHERE id = ? AND status = 0"
	db.Exec(sql, filename, time.Now().Unix(), data.ID)

	state.UnmarshalBinary([]byte(utils.Response200(200, "修改成功")))
	res.State = &state

	data.Euid = *Params.Body.Euid
	data.ID = 0
	data.Avatar = Params.Body.Avatar
	res.Data = &data

	o.Context.Respond(rw, r, route.Produces, route, res)

}
