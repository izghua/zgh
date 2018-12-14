/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-15
 * Time: 00:21
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
