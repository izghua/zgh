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
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *xorm.Engine


//type Mysql interface {
//	DbUser()
//	DbHost()
//	DbPort()
//	DbDatabase()
//	DbUserName()
//	DBPassword()
//}


//
//func (s *SqlParam)DbHost(host string) options {
//	return func(p *SqlParam) {
//		p.Host = host
//	}
//}
//
//func (s *SqlParam)DbPort(port string) func(*SqlParam) {
//	return func(p *SqlParam) {
//		p.Port = port
//	}
//}
//
//func (s *SqlParam)DbDataBase(dataBase string) func(*SqlParam) {
//	return func(p *SqlParam) {
//		p.DataBase = dataBase
//	}
//}
//
//func (s *SqlParam)DbUserName(userName string) func(*SqlParam) {
//	return func(p *SqlParam) {
//		p.UserName = userName
//	}
//}
//
//func (s *SqlParam)DbPassword(password string) func(*SqlParam) {
//	return func(p *SqlParam) {
//		p.Password = password
//	}
//}



type option func(*SqlParam) interface{}

func (p *SqlParam)DbUser(u string) option {
	return func(p *SqlParam) interface{} {
		user := p.User
		p.User = u
		return user
	}
}

func InitMysql(options ...func(*SqlParam)) *xorm.Engine {
	p := &SqlParam{}
	for _,option := range options {
		option(p)
	}

	fmt.Println(p,"看看具体实现了啥",options)
	dataSourceName := p.User + ":" + p.Password + "@/" + p.DataBase + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic("初始化数据库，创建连接异常:" + err.Error())
	}
	engine.TZLocation,_ = time.LoadLocation("Asia/Chongqing")
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
			fmt.Println(message)
			//go util.Alarm(message, util.ALARMALERT)
		}
		tryTimes = autoConnectMySQL(tryTimes, maxTryTimes)
	}
	return tryTimes
}

func MySQLAutoConnect() {
	autoConnectMySQL(0, 5)
}


