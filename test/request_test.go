/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-15
 * Time: 19:13
 */
package test

import (
	"fmt"
	"github.com/izghua/zgh/request"
	"github.com/izghua/zgh/utils"
	"izghua/zgh/conf"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	resp := new(utils.QqCaptchaResponse)
	res, _,err := request.New().Get(conf.QCapUrl).
		Param("aid","3333").
		Param("AppSecretKey","232342").
		Param("Ticket","23423").
		Param("Randstr","234324").
		Param("UserIP","127.0.0.1").
		Timeout(time.Minute*time.Duration(1)).Type(request.TypeUrlencoded).EndStruct(resp)
	fmt.Println(res,err)
}
