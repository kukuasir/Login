package utils

import (
	"encoding/json"
	"crypto/rand"
	"fmt"
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

func RandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
