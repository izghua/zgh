/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 21:58
 */
package utils

import (
	"net/smtp"
	"strings"
)

type EmailType string

type EmailParam struct {
	User EmailType `json:"user"`
	Password EmailType `json:"password"`
	Host EmailType `json:"host"`
	To EmailType `json:"to"`
	Subject EmailType `json:"subject"`
	Body EmailType `json:"body"`
}

var mailParam *EmailParam

var mailAddr,mailPort string

type EM  func(*EmailParam) interface{}

func (et EmailType) CheckIsNull() {
	if string(et) == "" {
		panic("不能为空")
	}
}

func (ep *EmailParam)SetMailUser(user EmailType) EM {
	return func(e *EmailParam) interface{} {
		u := e.User
		user.CheckIsNull()
		e.User = user
		return u
	}
}

func (ep *EmailParam)SetMailPwd(pwd EmailType) EM {
	return func(ep *EmailParam) interface{} {
		p := ep.Password
		pwd.CheckIsNull()
		ep.Password = pwd
		return p
	}
}

func (et EmailType)IsRight() {
	arr := strings.Split(string(et),":")
	if len(arr) != 2 {
		panic("有错误,可能不是分号")
	}
	mailAddr = arr[0]
	mailPort = arr[1]
}

func (ep *EmailParam)SetMailHost(host EmailType) EM {
	return func(ep *EmailParam) interface{} {
		h := ep.Host
		host.CheckIsNull()
		host.IsRight()
		ep.Host = host
		return h
	}
}

func (ep *EmailParam)MailInit(options ...EM) *EmailParam {
	q := &EmailParam{
	}
	for _,option := range options {
		option(q)
	}
	mailParam = q
	return mailParam
}


func SendMail( to, subject, body, mailType string) error {
	user := string(mailParam.User)
	password := string(mailParam.Password)
	host := string(mailParam.Host)
	auth := smtp.PlainAuth("", user, password, mailAddr)
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
		body = "<html><body>" + body + "</body></html>"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	msg = []byte(subject + contentType + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

