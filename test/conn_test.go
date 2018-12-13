/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-13
 * Time: 22:51
 */
package test

import (
	"github.com/izghua/zgh/conn"
	"testing"
)

func TestMysql(t *testing.T) {
	db := new(conn.Sp)
	dbUser := db.SetDbUserName("root")
	dbPwd := db.SetDbPassword("123456")
	dbPort := db.SetDbPort("3306")
	dbHost := db.SetDbHost("127.0.0.1")
	dbdb := db.SetDbDataBase("izghua")
	_,err := conn.InitMysql(dbUser,dbPwd,dbPort,dbHost,dbdb)
	if err != nil {
		t.Error("there is error")
	} else {
		t.Log("it is right")
	}
}


func TestRedis(t *testing.T) {
	rc := new(conn.RedisClient)
	addr := rc.SetRedisAddr("localhost:6379")
	pwd := rc.SetRedisPwd("")
	db := rc.SetRedisDb(0)
	_,err := rc.RedisInit(addr,db,pwd)
	if err != nil {
		t.Error("there is error")
	} else {
		t.Log("it is right")
	}
}