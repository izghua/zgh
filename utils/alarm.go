/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:29
 */
package utils

import (
	"fmt"
	"izghua/pkg/zgh/conf"
)

type AlarmParam struct {
	TypeOne string
}

var alarmParam *AlarmParam

type ap func(*AlarmParam) interface{}

func (alarm *AlarmParam)SetType(t string) ap {
	return func(alarm *AlarmParam) interface{} {
		ty := alarm.TypeOne
		alarm.TypeOne = t
		return ty
	}
}

type AlarmType string

const  (
	AlarmTypeOne AlarmType = "mail"
	AlarmTypeTwo AlarmType = "wechat"
	AlarmTypeThree AlarmType = "message"
)

func (at AlarmType)IsCurrentType() AlarmType {
	fmt.Println(at)
	return at
}

func (alarm *AlarmParam)Options(options ...ap) *AlarmParam {

	q := &AlarmParam{
		TypeOne:conf.ALARMTYPEONE,
	}
	for _,option := range options {
		option(q)
	}
	alarmParam = q
	return alarmParam
}



func Alarm(content string,priority string) {

	b := alarmParam.TypeOne
	fmt.Println(b,"222333")
}
