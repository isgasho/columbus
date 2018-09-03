package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

var engine *xorm.Engine

func openDB(username, password, addr, name string) *xorm.Engine {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	// TODO 打印sql连接日志
	engine, _ := xorm.NewEngine("mysql", config)

	setupDB(engine)

	return engine
}

func setupDB(engine *xorm.Engine) {
	// 设置db 连接属性
}

func InitSelfDB() *xorm.Engine {
	isDebug := viper.GetString("runmode")
	// Debug 模式先用sqlite3
	if isDebug == "debug"{
		engine, _ := xorm.NewEngine("sqlite3", "./test.db")
	    return engine
    }
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}