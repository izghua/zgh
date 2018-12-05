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

type Email struct {
	User string
	Password string
	Host string
	To string
	Subject string
	Body string
}

func SetMailUser() {

}



func SendMail(user, password, host, port, to, subject, body, mailType string) error {
	auth := smtp.PlainAuth("", user, password, port)
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

