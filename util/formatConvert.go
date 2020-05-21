package util

import (
	"github.com/gogf/gf/util/gconv"
)

//ToString 转换成字符串
func ToString(info interface{}) string {
	return gconv.String(info)
}

//ToInt 转换成int
func ToInt(info interface{}) int {
	return gconv.Int(info)
}
