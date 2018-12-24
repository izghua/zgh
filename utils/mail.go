/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-04
 * Time: 21:58
 */
package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"io/ioutil"
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
		zgh.ZLog().Error("content","value can not be null")
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
		zgh.ZLog().Error("may be is not semicolon")
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


func SendMail2(to string) error {

	//cc := "test@g9zz.com"
	sendTo := strings.Split(to, ";")

	subject := "这是一封演示用的邮件"
	boundary := "next message" //boundary 用于分割邮件内容，可自定义. 注意它的开始和结束格式

	mime := bytes.NewBuffer(nil)
	user := string(mailParam.User)
	password := string(mailParam.Password)
	host := string(mailParam.Host)
	//auth := smtp.PlainAuth("", user, password, mailAddr)
	//设置邮件
	mime.WriteString(fmt.Sprintf("From: %s<%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\n", user, user, to,  subject))
	//msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", boundary))
	mime.WriteString("Content-Description: 这是一封带附档的邮件\r\n")

	//邮件普通Text正文
	mime.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	mime.WriteString("This is a multipart message in MIME format.")

	//邮件HTML正文
	mime.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))

	boundaryHtml := "boundaryHtml"
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=%s\r\n", boundaryHtml))
	mime.WriteString("Content-Description: Message in alternative text and HTML forms\r\n")
	mime.WriteString(fmt.Sprintf("\n--%s\r\n", boundaryHtml))
	mime.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\r\n", "text/html"))
	mime.WriteString(`<html><body>
		<p><img src="https://golang.org/doc/gopher/doc.png"></p><br/>
		<h1>最近有点消沉,也有点想家了.</h1>
		</body></html>`)
	mime.WriteString(fmt.Sprintf("\n--%s--\r\n\r\n", boundaryHtml))

	// 第一个附件
	attaFile := "./backup"
	attaFileName := "backup.zip"
	mime.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))
	mime.WriteString("Content-Type: application/octet-stream\r\n")
	mime.WriteString("Content-Description: 附一个Go文件\r\n")
	mime.WriteString("Content-Transfer-Encoding: base64\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\"" + attaFileName + "\"\r\n\r\n")
	//读取并编码文件内容
	attaData, err := ioutil.ReadFile(attaFile)
	if err != nil {
		return err
	}
	b := make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
	base64.StdEncoding.Encode(b, attaData)
	mime.Write(b)

	//第二个附件
	mime.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	mime.WriteString("Content-Type: text/plain\r\n")
	mime.WriteString("Content-Description: 附一个Text文件\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\"test.txt\"\r\n\r\n")
	mime.WriteString("this is the attachment text")

	//邮件结束
	mime.WriteString("\r\n--" + boundary + "--\r\n\r\n")

	fmt.Println(mime.String())

	//发送相关
	//smtpHost, _, err := net.SplitHostPort(host)
	//if err != nil {
	//	return err
	//}
	auth := smtp.PlainAuth("", user, password, mailAddr)
	//err := smtp.SendMail(host, auth, user, sendTo, msg)

	return smtp.SendMail(host, auth, user, sendTo, mime.Bytes())

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

