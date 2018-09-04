package db

import (
	//"database/sql"
	//"errors"
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

var dns string

func Engine() (*xorm.Engine, error){
	initEngine()
    fmt.Printf("%s-------\n", dns)
    return xorm.NewEngine("mysql", dns)
}

func SqliteEngine() (*xorm.Engine, error) {
	f := "columbus.db"
	//os.Remove(f)

	return xorm.NewEngine("sqlite3", f)
}

func initEngine() {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.charset"))
}