/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:29
 */
package utils

import (
	"strings"
)

// Define AlarmType to string
// for to check the params is right
type AlarmType string

// this are some const params what i defined
// only this can be to input
const  (
	AlarmTypeOne AlarmType = "mail"
	AlarmTypeTwo AlarmType = "wechat"
	AlarmTypeThree AlarmType = "message"
)

type AlarmParam struct {
	Types AlarmType
}

var alarmParam *AlarmParam

// Define a closure type to next
type ap func(*AlarmParam) interface{}

// can use this function to set a new value
// but to check it is a right type
func (alarm *AlarmParam)SetType(t AlarmType) ap {
	return func(alarm *AlarmParam) interface{} {
		str := strings.Split(string(t),",")
		if len(str) == 0 {
			panic("有错误,不能不传入任何值")
		}
		for _,types := range str {
			s := AlarmType(types)
			s.IsCurrentType()
		}
		ty := alarm.Types
		alarm.Types = t
		return ty
	}
}

// judge it is a right type what i need
// if is it a wrong type, i must return a panic to above
func (at AlarmType)IsCurrentType() AlarmType {
	switch at {
	case AlarmTypeOne:
		return at
	case AlarmTypeTwo:
		return at
	case AlarmTypeThree:
		return at
	default:
		panic("有错误")
	}

	return at
}

// implementation value
func (alarm *AlarmParam)AlarmInit(options ...ap) *AlarmParam {
	q := &AlarmParam{
	}
	for _,option := range options {
		option(q)
	}
	alarmParam = q
	return alarmParam
}



func Alarm(content string,priority string) {


}
