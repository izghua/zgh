/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 21:35
 */
package utils

import "time"

var (
	TimeLocation, _ = time.LoadLocation("Asia/Chongqing") //当地时间
)

// 返回当前时间格式
func GetDateTime() string {
	return time.Now().In(TimeLocation).Format("2006-01-02 15:04:05")
}

