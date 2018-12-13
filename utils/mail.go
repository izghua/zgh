/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 21:58
 */
package utils

import (
	"errors"
	"izghua/pkg/zgh/conf"
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
	MailType EmailType `json:"mail_type"`
}

var mailParam *EmailParam

var mailAddr string

type EM  func(*EmailParam) (interface{},error)

func (et EmailType) CheckIsNull() error {
	if string(et) == "" {
		ZLog().Error("content","value can not be null")
		return errors.New("value can not be null")
	}
	return nil
}

func (ep *EmailParam)SetMailUser(user EmailType) EM {
	return func(e *EmailParam) (interface{},error) {
		u := e.User
		err := user.CheckIsNull()
		if err != nil {
			return nil,err
		}
		e.User = user
		return u,nil
	}
}

func (ep *EmailParam)SetMailPwd(pwd EmailType) EM {
	return func(ep *EmailParam) (interface{},error) {
		p := ep.Password
		err := pwd.CheckIsNull()
		if err != nil {
			return nil,err
		}
		ep.Password = pwd
		return p,nil
	}
}

func (et EmailType)IsRight() error {
	arr := strings.Split(string(et),":")
	if len(arr) != 2 {
		ZLog().Error("may be is not semicolon")
		return errors.New("may be is not semicolon")
	}
	mailAddr = arr[0]
	return nil
}

func (ep *EmailParam)SetMailHost(host EmailType) EM {
	return func(ep *EmailParam) (interface{},error) {
		h := ep.Host
		err := host.CheckIsNull()
		if err != nil {
			return nil,err
		}
		err = host.IsRight()
		if err != nil {
			return nil,err
		}
		ep.Host = host
		return h,nil
	}
}

func (ep *EmailParam)SetMailType(types EmailType) EM {
	return func(ep *EmailParam) (interface{},error) {
		ty := ep.MailType
		err := types.CheckIsNull()
		if err != nil {
			return nil,err
		}
		ep.MailType = ty
		return ty,nil
	}
}


func (ep *EmailParam)MailInit(options ...EM) error {
	q := &EmailParam{
		MailType:conf.MAIlTYPE,
	}
	for _,option := range options {
		_,err := option(q)
		if err != nil {
			return err
		}
	}
	mailParam = q
	return nil
}


func SendMail( to string, subject string, body string) error {
	user := string(mailParam.User)
	password := string(mailParam.Password)
	host := string(mailParam.Host)
	auth := smtp.PlainAuth("", user, password, mailAddr)
	var contentType string
	if mailParam.MailType == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
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

