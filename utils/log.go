/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 21:33
 */
package utils

import (
	"errors"
	"fmt"
	"github.com/Penglq/QLog"
	"github.com/izghua/zgh/conf"
	"log"
	"runtime"
)

type ZLogParam struct {
	FilePath string
	FileName string
	FileSuffix string
	FileMaxSize int64
	FileMaxNSize int
	TimeZone string
}

var zLogParam *ZLogParam

type zp func(*ZLogParam) interface{}

func (zlp *ZLogParam) SetFilePath(fp string) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.FilePath
		zlp.FilePath = fp
		return f
	}
}

func (zlp *ZLogParam) SetFileName(fn string) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.FileName
		zlp.FileName = fn
		return f
	}
}

func (zlp *ZLogParam) SetFileSuffix(fs string) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.FileSuffix
		zlp.FileSuffix = fs
		return f
	}
}

func (zlp *ZLogParam) SetFileMaxSize(fms int64) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.FileMaxSize
		zlp.FileMaxSize = fms
		return f
	}
}


func (zlp *ZLogParam) SetFileMaxNSize(fmns int) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.FileMaxNSize
		zlp.FileMaxNSize = fmns
		return f
	}
}

func (zlp *ZLogParam) SetTimeZone(tz string) zp {
	return func(zlp *ZLogParam) interface{} {
		f := zlp.TimeZone
		zlp.TimeZone = tz
		return f
	}
}

var Zog QLog.LoggerInterface

func (zlp *ZLogParam)ZLogInit(options ...zp) error {
	q := &ZLogParam{
		FilePath:conf.LOGFILEPATH,
		FileName:conf.LOGFILENAME,
		FileSuffix:conf.LOGFILESUFFIX,
		FileMaxSize:conf.LOGFILEMAXSIZE,
		FileMaxNSize:conf.LOGFILEMAXNSIZE,
		TimeZone:conf.LOGTIMEZONE,
	}
	for _,option := range options {
		option(q)
	}
	zLogParam = q
	if zLogParam == nil {
		log.Fatalf("Zlog not init err %s", errors.New("日志没有初始化 - "))
	}
	l := QLog.GetLogger()
	l.SetConfig(QLog.INFO, zLogParam.TimeZone,
		QLog.WithFileOPT(zLogParam.FilePath, zLogParam.FileName, zLogParam.FileSuffix, zLogParam.FileMaxSize,zLogParam.FileMaxNSize),
		QLog.WithConsoleOPT(),
	)

	Zog = l
	return nil
}

// the log is designed by my colleague
// https://github.com/Penglq/QLog
// i just package it
// you must input content what it is wrong content
// then you must describe it is type
func ZLog() QLog.LoggerInterface {
	funcName,file,_,ok := runtime.Caller(1)
	fmt.Println(ok,"日志是否正确",funcName,file)
	if ok {
		fName := runtime.FuncForPC(funcName).Name()
		fmt.Println("隔断",fName,"看日志陆军")
		Zog.AddTextPrefix("method",fName)
	}
	return Zog
}

