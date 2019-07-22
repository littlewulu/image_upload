package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
)

const (
	randStrSource = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)


// 获取随机字符串
func RanString(l int)string{
	length := len(randStrSource)
	r := make([]byte, l)
	for i := 0; i < l; i++{
		r[i] = randStrSource[rand.Intn(length)]
	}
	return string(r)
}

// 计算MD5
func CalMd5(data []byte)string{
	return fmt.Sprintf("%x", md5.Sum(data))
}

// 检测文件类型
func DetectContentType(buff []byte)(string, error){
	size := 512
	if len(buff) < size{
		size = len(buff)
	}
	filetype := http.DetectContentType(buff[0:size])
	return filetype, nil
}

// 校验文件类型合法性
func ValidateImageType(t string)(string, bool){
	switch t {
	case "image/jpeg", "image/jpg":
		return "jpg", true
	case "image/png":
		return "png", true
	case "image/gif":
		return "gif", true
	default:
		return "", false
	}
}
