package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"sync"
)

var globalOrm orm.Ormer
var once sync.Once

func init() {
	dbUrl := beego.AppConfig.String("mysql_url")
	database := beego.AppConfig.String("mysql_database")
	dbUser := beego.AppConfig.String("mysql_user")
	dbPasswd := beego.AppConfig.String("mysql_pwd")
	dbPort := beego.AppConfig.String("mysql_port")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPasswd, dbUrl, dbPort, database) +
		"&loc=" + url.QueryEscape("Asia/Shanghai")
	logs.Info("sql connection url : %s", connStr)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connStr, 30)
	orm.RegisterModelWithPrefix("api_", new(Account), new(Auth), new(Edge), new(Role))
	err := orm.RunSyncdb("default", false, true)
	orm.Debug = true
	if err != nil {
		logs.Info("RegisterModel err : %r", err)
		panic(err)
	}
}

// global singleton orm instance
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
