/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-13
 * Time: 22:51
 */
package test

import (
	"github.com/izghua/zgh/utils/hashid"
	"github.com/izghua/zgh/utils/mail"
	"testing"
)

// the alarm init must be the zlog init first
// so ...
//func TestAlarm(t *testing.T) {
//	zlog := new(zgh.ZLogParam)
//	err :=zlog.ZLogInit()
//	if err != nil {
//		t.Error("it is err")
//	} else {
//		t.Log("it is right")
//	}
//	alarm := new(utils.AlarmParam)
//	alarmT := alarm.SetType("mail")
//	mailTo := alarm.SetMailTo("xzghua@gmail.com")
//	err = alarm.AlarmInit(alarmT,mailTo)
//	if err != nil {
//		t.Error("it is err")
//	} else {
//		t.Log("it is right")
//	}
//}



func TestMail(t *testing.T) {
	mail := new(mail.EmailParam)
	mailUser := mail.SetMailUser("test@test.com")
	mailPwd := mail.SetMailPwd("test")
	mailHost :=  mail.SetMailHost("smtp.mxhichina.com:25")
	_,err := mail.MailInit(mailPwd,mailHost,mailUser)
	if err != nil {
		t.Error("it is err MailInit")
	} else {
		t.Log("it is right MailInit")
	}
}

func TestHashId(t *testing.T) {
	hd := new(hashid.HashIdParams)
	salt := hd.SetHashIdSalt("test salt")
	hdLength := hd.SetHashIdLength(10)
	_,err := hd.HashIdInit(hdLength,salt)
	if err != nil {
		t.Error("it is err")
	} else {
		t.Log("it is right")
	}
}

//func TestZLog(t *testing.T) {
//	zlog := new(zgh.ZLogParam)
//	zlog.FilePath = "./log"
//	zlog.FileSuffix = "zog"
//	zlog.FileName = "zlog"
//	zlog.FileMaxNSize = 1
//	zlog.FileMaxSize = 0
//	err :=zlog.ZLogInit()
//	if err != nil {
//		t.Error("it is err")
//	} else {
//		t.Log("it is right")
//	}
//}