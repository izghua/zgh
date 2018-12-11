/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 22:29
 */
package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Define AlarmType to string
// for to check the params is right
type AlarmType string

type AlarmMailReceive string

// this are some const params what i defined
// only this can be to input
const  (
	AlarmTypeOne AlarmType = "mail"
	AlarmTypeTwo AlarmType = "wechat"
	AlarmTypeThree AlarmType = "message"
	//AlarmMailTo AlarmMailReceive = ""
)

type AlarmParam struct {
	Types AlarmType
	MailTo AlarmMailReceive
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

func (alarm *AlarmParam)SetMailTo(t AlarmMailReceive) ap {
	return func(alarm *AlarmParam) interface{} {
		to := alarm.MailTo
		t.CheckIsNull().MustMailFormat()
		alarm.MailTo = t
		return to
	}
}

// alarm receive account can not null
func (t AlarmMailReceive)CheckIsNull() AlarmMailReceive {
	if len(t) == 0 {
		panic("不能为空")
	}
	return t
}

// alarm receive account must be mail format
func (t AlarmMailReceive)MustMailFormat() AlarmMailReceive {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", string(t)); !m {
		panic("格式不正确")
	}
	return t
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
	fmt.Println(alarmParam.MailTo,alarmParam.MailTo == "")
	types := strings.Split(string(alarmParam.Types),",")
	var err error
	for _,a := range types {
		switch AlarmType(a) {
		case AlarmTypeOne:
			if alarmParam.MailTo == "" {
				panic("邮件接收者不能为空")
			}
			err = SendMail(string(alarmParam.MailTo),"报警",content)
			break
		case AlarmTypeTwo:
			break
		case AlarmTypeThree:
			break
		}
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println(alarmParam.Types,"baojing",alarmParam.MailTo)
}
