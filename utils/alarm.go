/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:29
 */
package utils

import "izghua/pkg/zgh/conf"

type AlarmParam struct {
	TypeOne string
}

type ap func(*AlarmParam) interface{}

func (alarm *AlarmParam)SetType(t string) ap {
	return func(alarm *AlarmParam) interface{} {
		ty := alarm.TypeOne
		alarm.TypeOne = t
		return ty
	}
}

func (alarm *AlarmParam)Options(options ...ap) *AlarmParam {
	q := &AlarmParam{
		TypeOne:conf.ALARMTYPE1,
	}
	for _,option := range options {
		option(q)
	}
	return q
}


func Alarm(content string) {

}
