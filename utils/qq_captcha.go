/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-15
 * Time: 16:44
 */
package utils

import (
	"fmt"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/request"
	"time"
)

type QQCaptcha struct {
	Aid string
	AppSecretKey string
	Ticket string
	Randstr string
	UserIP  string
	Url string
}


type qct func(qc *QQCaptcha) interface{}

func (qc *QQCaptcha) SetAid(aid string) qct {
	return func(qc *QQCaptcha) interface{} {
		a := qc.Aid
		qc.Aid = aid
		return a
	}
}

func (qc *QQCaptcha) SetSecretKey(sk string) qct {
	return func(qc *QQCaptcha) interface{} {
		a := qc.AppSecretKey
		qc.AppSecretKey = sk
		return a
	}
}

var qqCaptcha *QQCaptcha


func (qc *QQCaptcha)QQCaptchaInit(options ...qct) error {
	q := &QQCaptcha{
	}
	for _,option := range options {
		option(q)
	}
	qqCaptcha = q
	return nil
}

type QqCaptchaResponse struct {
	Response int `json:"response"`
	EvilLevel int `json:"evil_level"`
	errMsg string `json:"err_msg"`
}

func QQCaptchaVerify(ticket string,randStr string,userIP string) {
	resp := new(QqCaptchaResponse)
	res, _,err := request.New().Get(conf.QCapUrl).
		Param("aid",qqCaptcha.Aid).
		Param("AppSecretKey",qqCaptcha.AppSecretKey).
		Param("Ticket",ticket).
		Param("Randstr",randStr).
		Param("UserIP",userIP).
		Timeout(time.Minute*time.Duration(1)).Type(request.TypeUrlencoded).EndStruct(resp)
	fmt.Println(res,err)
}