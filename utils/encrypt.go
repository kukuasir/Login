package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encrypt(text string) string {
	hashstr := md5.New()
	hashstr.Write([]byte(text))
	return hex.EncodeToString(hashstr.Sum(nil))
}

func GenNickNameBy(phone string) string {
	if len(phone) == 11 {
		return phone[0:4] + "****" + phone[7:11]
	} else if len(phone) > 0 {
		return phone
	} else {
		return ""
	}
}
