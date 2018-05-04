package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
)

func MD5Encrypt(text string) string {
	hashstr := md5.New()
	hashstr.Write([]byte(text))
	return hex.EncodeToString(hashstr.Sum(nil))
}

func GenEuidBy(uname string) string {
	str := "BTK_" + uname + strconv.FormatInt(time.Now().Unix(), 10)
	return MD5Encrypt(str)
}

func GenNickNameBy(phone string) string {
	return phone[0:4] + "****" + phone[7:11]
}
