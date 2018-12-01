/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-11-29
 * Time: 23:42
 */
package conn

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"gitlab.yixinonline.org/pkg/bang/conf"
	"gitlab.yixinonline.org/pkg/bang/util"
	"time"
)

var mysql *xorm.Engine

type SqlParam struct {
	User string
	Host string
	Port string
	DataBase string
	UserName string
	Password string
}

func InitMysql(s *SqlParam) *xorm.Engine {
	dataSourceName := s.User + ":" + s.Password + "@/" + s.DataBase + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic("初始化数据库，创建连接异常:" + err.Error())
	}
	engine.TZLocation = conf.TimeLocation
	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)
	engine.SetConnMaxLifetime(0)
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	mysql = engine
	timer := time.NewTicker(time.Minute * 10)
	go func(conn *xorm.Engine) {
		for _ = range timer.C {
			if err = mysql.Ping(); err != nil {
				MySQLAutoConnect()
			}
		}
	}(mysql)
	return mysql
}

func autoConnectMySQL(tryTimes int, maxTryTimes int) int {
	tryTimes++
	if tryTimes <= maxTryTimes {
		if mysql.Ping() != nil {
			message := fmt.Sprintf("数据库连接失败,已重连%d次", tryTimes)
			//yrdLog.GetLogger().Error("mysql", message)
			go util.Alarm(message, util.ALARMALERT)
		}
		tryTimes = autoConnectMySQL(tryTimes, maxTryTimes)
	}
	return tryTimes
}

func MySQLAutoConnect() {
	autoConnectMySQL(0, 5)
}


