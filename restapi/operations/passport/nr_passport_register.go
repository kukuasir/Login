// Code generated by go-swagger; DO NOT EDIT.

package passport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"Passport/models"
	"Passport/utils"
	"fmt"
	"time"
)

// NrPassportRegisterHandlerFunc turns a function with the right signature into a passport register handler
type NrPassportRegisterHandlerFunc func(NrPassportRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrPassportRegisterHandlerFunc) Handle(params NrPassportRegisterParams) middleware.Responder {
	return fn(params)
}

// NrPassportRegisterHandler interface for that can handle valid passport register params
type NrPassportRegisterHandler interface {
	Handle(NrPassportRegisterParams) middleware.Responder
}

// NewNrPassportRegister creates a new http.Handler for the passport register operation
func NewNrPassportRegister(ctx *middleware.Context, handler NrPassportRegisterHandler) *NrPassportRegister {
	return &NrPassportRegister{Context: ctx, Handler: handler}
}

/*NrPassportRegister swagger:route POST /passport/register Passport passportRegister

注册接口

注册接口

*/
type NrPassportRegister struct {
	Context *middleware.Context
	Handler NrPassportRegisterHandler
}

func (o *NrPassportRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrPassportRegisterParams()

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

	var res models.RegisterState
	var state models.State
	var user models.UserBase

	// 定义错误信息
	var code int64
	var message string

	// 先判断验证码是否正确
	var sms SMSRecord
	db.Table(utils.T_SMS).Where("phone=?", *Params.Body.Phone).Where("code=?", *Params.Body.ValidCode).Where("type=0").Order("create_at DESC").First(&sms)
	if len(sms.Code) == 0 {
		code = 401
		message = "验证码不正确"
	} else {
		if time.Now().Unix() - utils.T_EXPIRED_SECONDS > sms.CreateAt {
			code = 402
			message = "验证码已失效"
		} else {

			QueryUser(db, *Params.Body.Phone, &user)

			// 用户ID不存在
			if user.Euid == nil {
				// 检查用户昵称是否被占用
				var tmp models.UserBase
				db.Table(utils.T_USER).Where("nick_name=?", Params.Body.NickName).Find(&tmp)
				if len(tmp.NickName) == 0 {

					// 如果没有NickName，默认使用手机号作为昵称
					nick_name := Params.Body.NickName
					if len(Params.Body.NickName) == 0 {
						nick_name = utils.GenNickNameBy(*Params.Body.Phone)
						fmt.Println("nick_name = ", nick_name)
					}
					sql := "INSERT INTO btk_User(euid,nick_name,birth_day,gender,phone,password,invite_code,register_at) VALUES(?,?,?,?,?,?,?,?)"
					db.Exec(sql, utils.GenEuidBy(*Params.Body.Phone), nick_name, Params.Body.BirthDay, Params.Body.Gender, Params.Body.Phone, utils.MD5Encrypt(*Params.Body.Password), Params.Body.InviteCode, time.Now().Unix())
					// 注册成功后，再去查一下
					QueryUser(db, *Params.Body.Phone, &user)
					// 头像路径加上域名
					if len(user.Avatar) > 0 {
						user.Avatar = utils.T_IMAGE_DOMAIN + user.Avatar
					}
					res.Data = &user
					code = 200
					message = "注册成功"
				} else {
					code = 202
					message = "用户名已被注册，请更换用户名"
				}
			} else {
				code = 203
				message = "手机号已被注册"
			}
		}
	}

	state.UnmarshalBinary([]byte(utils.Response200(code, message)))
	res.State = &state

	o.Context.Respond(rw, r, route.Produces, route, res)

}

func QueryUser(db *gorm.DB, uname string, user *models.UserBase) {
	sql := "SELECT euid,nick_name,avatar,phone,birth_day,current_coins,current_points,friends_count,gender,level,login_at,register_at FROM btk_User WHERE status = 0 && phone = " + uname
	db.Raw(sql).Find(user)
}
