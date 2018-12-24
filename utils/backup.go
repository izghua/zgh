/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-22
 * Time: 00:16
 */
package utils

import (
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/conn"
	"os"
	"time"
)

type BackUpParam struct {
	Files []*os.File
	Dest string
	Duration time.Duration
	FileName string
	FilePath string
}

func (bp *BackUpParam) SetFiles(f []*os.File) *BackUpParam {
	bp.Files = f
	return bp
}

func (bp *BackUpParam) SetDest(d string) *BackUpParam {
	bp.Dest = d
	return bp
}

func (bp *BackUpParam) SetDuration(d time.Duration) *BackUpParam {
	bp.Duration = d
	return bp
}

func (bp *BackUpParam) SetFileName(fn string) *BackUpParam {
	bp.FileName = fn
	return bp
}

func (bp *BackUpParam) SetFilePath(fp string) *BackUpParam {
	bp.FilePath = fp
	return bp
}

var backUp *BackUpParam

func BackUpInit() {
	q := &BackUpParam{
		Dest:conf.BackUpDest,
		Duration:conf.BackUpDuration,
		FileName:time.Now().Format("2006-01-02") + conf.BackUpSqlFileName,
		FilePath:conf.BackUpFilePath,
	}
	backUp = q
}


func (bp *BackUpParam)Backup() error {
	backUp = bp

	err := conn.SqlDump(backUp.FileName,backUp.FilePath)
	if err != nil {
		zgh.ZLog().Error("message","back up sql dump is error","error",err.Error())
		return err
	}

	err = Compress(backUp.Files,backUp.Dest)
	if err != nil {
		zgh.ZLog().Error("message","back up compress is error","error",err.Error())
		return err
	}

	return nil


}
