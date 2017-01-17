package models

import (
	_ "github.com/astaxie/beego/validation"
	"time"
)

type Auth struct {
	Id         uint64    `orm:"auto" valid:"Required"`
	Name       string    `orm:"size(50)" valid:"Required"`
	ServerDns  string    `orm:"size(50)" valid:"Required"`
	IP         string    `orm:"size(20)" valid:"Required;IP"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func InsertorUpdateAuth(auth Auth) {
	db := GetOrmer()
	if existed := db.QueryTable("api_auth").Filter("id", auth.Id).Exist(); existed == false {
		db.Insert(&auth)
	} else {
		db.Update(&auth)
	}
}
