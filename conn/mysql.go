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
	"izghua/pkg/zgh/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *xorm.Engine

type SqlParam struct {
	Host string
	Port string
	DataBase string
	UserName string
	Password string
}

type sp  func(*SqlParam) interface{}

func (p *SqlParam)SetDbHost(host string) sp {
	return func(p *SqlParam) interface{} {
		h := p.Host
		p.Host = host
		return h
	}
}

func (p *SqlParam)SetDbPort(port string) sp {
	return func(p *SqlParam) interface{} {
		pt := p.Port
		p.Port = port
		return pt
	}
}

func (p *SqlParam)SetDbDataBase(dataBase string) sp {
	return func(p *SqlParam) interface{} {
		db := p.DataBase
		p.DataBase = dataBase
		return db
	}
}


func (p *SqlParam)SetDbPassword(pwd string) sp {
	return func(p *SqlParam) interface{} {
		password := p.Password
		p.Password = pwd
		return password
	}
}


func (p *SqlParam)SetDbUserName(u string) sp {
	return func(p *SqlParam) interface{} {
		user := p.UserName
		p.UserName = u
		return user
	}
}

func InitMysql(options ...sp) *xorm.Engine {
	q := &SqlParam{
		Host:conf.DBHOST,
		Port:conf.DBPORT,
		Password:conf.DBPASSWORD,
		DataBase:conf.DBDATABASE,
		UserName:conf.DBUSERNAME,
	}
	for _,option := range options {
		option(q)
	}

	dataSourceName := q.UserName + ":" + q.Password + "@/" + q.DataBase + "?charset=utf8"
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


