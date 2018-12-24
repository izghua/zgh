/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-22
 * Time: 00:16
 */
package backup

import (
	"github.com/izghua/zgh"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/conn"
	"github.com/izghua/zgh/utils/cron"
	"github.com/izghua/zgh/utils/zip"
	"os"
	"time"
)




func (bp *BackUpParam) SetFiles(f []*os.File) *BackUpParam  {
	bp.Files = f
	return bp
}

func (bp *BackUpParam) SetDest(d string) *BackUpParam  {
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

type BackUpParam struct {
	Files  []*os.File `json:"files"`
	Duration time.Duration `json:"duration"`
	Dest string `json:"dest"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}


func (bp *BackUpParam)FilePathIsNull() *BackUpParam {
	if bp.FilePath == "" {
		zgh.ZLog().Warn("message","data is null")
		bp.SetFilePath(conf.BackUpFilePath)
	}
	return bp
}

func (bp *BackUpParam)DestIsNull() *BackUpParam {
	if bp.Dest == "" {
		zgh.ZLog().Warn("message","data is null")
		bp.SetDest(conf.BackUpDest)
	}
	return bp
}

func (bp *BackUpParam)FileNameIsNull() *BackUpParam {
	if bp.FileName == "" {
		zgh.ZLog().Warn("message","data is null")
		bp.SetFileName(time.Now().Format("2006-01-02") + conf.BackUpSqlFileName)
	}
	return bp
}

func (bp *BackUpParam)DurationIsNull() *BackUpParam {
	if bp.Duration == 0 {
		zgh.ZLog().Warn("message","data is null")
		bp.SetDuration(conf.BackUpDuration)
	}
	return bp
}



func (bp *BackUpParam)Backup() error {
	bp.DestIsNull().FileNameIsNull().FilePathIsNull().DurationIsNull()
	//backUp = bp
	//fmt.Println("文件名filename",backUp.FileName)
	//fmt.Println("文件位置FilePath",backUp.FilePath)
	//fmt.Println("备份循环时间Duration",backUp.Duration)
	//fmt.Println("目标目录Dest",backUp.Dest)
	//
	////fmt.Println(backUp.FileName,backUp.FilePath,"看问题",backUp.Duration)
	cron.ZgCron(bp.Duration,bp.doBackUp)

	return nil
}

func (bp *BackUpParam)doBackUp() {
	err := conn.SqlDump(bp.FileName,bp.FilePath)
	if err != nil {
		zgh.ZLog().Error("message","back up sql dump is error","error",err.Error())
	}

	err = zip.Compress(bp.Files,bp.Dest)
	if err != nil {
		zgh.ZLog().Error("message","back up compress is error","error",err.Error())
		return
	}
	return
}