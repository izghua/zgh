/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-03
 * Time: 21:33
 */
package utils

import (
	"fmt"
	"github.com/Penglq/QLog"
	"izghua/pkg/zgh/conf"
	"time"
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

func (zlp *ZLogParam)ZogInit(options ...zp) *ZLogParam {
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
	return zLogParam
}


// the log is designed by my colleague
// https://github.com/Penglq/QLog
// i just package it
// you must input content what it is wrong content
// then you must describe it is type
func ZLog(content string ,priority string) {
	fmt.Println(zLogParam,"看日志",time.Now().Format(time.RFC3339))
	if zLogParam == nil {
		panic("日志木有初始化")
	}
	l := QLog.GetLogger()
	l.SetConfig(QLog.INFO, zLogParam.TimeZone,
		QLog.WithFileOPT(zLogParam.FilePath, zLogParam.FileName, zLogParam.FileSuffix, zLogParam.FileMaxSize,zLogParam.FileMaxNSize),
		QLog.WithConsoleOPT(),
	)
	l.Info("a","有错","b","还是有错","content",content)
}

