package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//	每个事件生成唯一guid,标识事件的唯一性，便于发送模块对数据库事件的更新修改操作
func GetMD5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return GetMD5String(base64.URLEncoding.EncodeToString(b))
}
