package database

import (
	//"database/sql"
	//"errors"
	"fmt"

	"columbus/config"

	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/pflag"
)

var DB *xorm.Engine

var dns string

var err error

var (
	cfg = pflag.StringP("config", "c", "", "columbus config file path.")
)

func init() {
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 启动时就打开数据库连接
	if err = initEngine(); err != nil {
		panic(err)
	}
}

func initEngine() error {
	dbEngine := viper.GetString("db.engine")
	if dbEngine == "postgre" {
		DB, err = postgreEngine()
	} else if dbEngine == "mysql" {
		DB, err = mysqlEngine()
	} else {
		DB, err = sqliteEngine()
	}
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(viper.GetInt("db.max_conn"))
	DB.ShowSQL(viper.GetBool("db.show_sql"))

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

	return nil
}

func mysqlEngine() (*xorm.Engine, error){
	setEngine()
    fmt.Printf("mysql db config-> %s\n", dns)
    return xorm.NewEngine("mysql", dns)
}

func postgreEngine() (*xorm.Engine, error){
	setEngine()
	fmt.Printf("postgre db config-> %s\n", dns)
	return xorm.NewEngine("postgres", dns)
}

func sqliteEngine() (*xorm.Engine, error) {
	f := "columbus.db"
	//os.Remove(f)
	fmt.Printf("sqlite db")
	return xorm.NewEngine("sqlite3", f)
}

func setEngine() {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.charset"))
}