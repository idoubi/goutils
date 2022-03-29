package goutils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// MD5 md5加密
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))

	return hex.EncodeToString(ctx.Sum(nil))
}

// Base64Encode base64加密
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode base64解密
func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}

	return string(data)
}

// SortEncodeMap 自定义编码 map[string]string
func SortEncodeMap(m map[string]string, key, secret string) string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	// 字典序排列
	sort.Strings(keys)

	arr := []string{}
	str := ""

	for _, k := range keys {
		if m[k] == "" {
			continue
		}
		arr = append(arr, fmt.Sprintf("%s=%s", k, m[k]))
	}

	// 拼接参数
	str = strings.Join(arr, "&")

	if secret != "" {
		// 拼接秘钥
		if key == "" {
			str += secret
		} else {
			str += fmt.Sprintf("&%s=%s", key, secret)
		}
	}

	return str
}

// SortEncodeMapValue 编码map数值
func SortEncodeMapValue(m map[string]string) string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	// 字典序排列
	sort.Strings(keys)

	arr := []string{}
	str := ""

	for _, k := range keys {
		arr = append(arr, fmt.Sprintf("%s", m[k]))
	}

	// 拼接参数
	str = strings.Join(arr, "")

	return str
}

// Sha256SignByKey Sha256方式签名source
func Sha256SignByKey(key string, source string) string {
	res := sha256.Sum256([]byte(source + "&key=" + key))

	return hex.EncodeToString(res[:])
}

// HmacSha256 加密
func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha256WithBase64 hmac-sha256后使用base64编码
func HmacSha256WithBase64(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
