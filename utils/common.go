package utils

import (
	"encoding/json"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/** 所需的表名定义 */
const (
	T_USER = "btk_User"
	T_SMS = "btk_SMS"
)

/** 短信验证码失效时间 */
const T_EXPIRED_SECONDS = 300

/** 定义头像图片域名 */
const T_IMAGE_DOMAIN = "http://inj-zone-img.bitekun.xin/resource/avatar/"

/** 定义快捷登录Platform */
const T_PLATFORM_QUICK_LOGIN = 99

/** 定义返回成功后的方法 */
type Response struct {
	 Code int64 `json:"code"`
	 Message string `json:"message"`
}
func Response200(code int64, msg string) string {
	var response Response
	response.Code = code
	response.Message = msg
	out, _ := json.Marshal(response)
	return string(out)
}

/**
生成随机验证码
*/
func RandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b) + strconv.FormatInt(time.Now().Unix(), 10)
}

/**
加密用户ID
*/
func EncryptEuid(uid int64) string {
	return strconv.FormatInt(uid, 10)
}

/**
解密用户ID
*/
func DecodeUserID(euid string) int64 {
	userId, _ := strconv.ParseInt(euid, 10, 64)
	return userId
}

/**
补充图片域名
第三方登录的产生的图片不需要增加域名
*/
func CompleteImage(name string) string {
	if strings.Contains(name, "http") || strings.Contains(name, "https") {
		return name
	} else {
		return T_IMAGE_DOMAIN + name
	}
}
