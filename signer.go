package tsgutils

/*
 sign utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 signing a message
 using: hmac sha256 + base64
  eg:
    message = Pre_hash function comment
    secretKey = E65791902180E9EF4510DB6A77F6EBAE

  return signed string = TO6uwdqz+31SIPkd4I+9NiZGmVH74dXi+Fd5X0EzzSQ=
*/
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

/*
 the pre hash string
  eg:
    timestamp = 2018-03-08T10:59:25.789Z
    method  = POST
    request_path = /orders?before=2&limit=30
    body = {"product_id":"BTC-USD-0309","order_id":"377454671037440"}

  return pre hash string = 2018-03-08T10:59:25.789ZPOST/orders?before=2&limit=30{"product_id":"BTC-USD-0309","order_id":"377454671037440"}
*/
func PreHashString(timestamp string, method string, requestPath string, body string) string {
	return timestamp + strings.ToUpper(method) + requestPath + body
}

/*
  md5 sign: "123" -> "202cb962ac59075b964b07152d234b70"
*/
func Md5Signer(message string) string {
	data := []byte(message)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
