package models

import (
	"time"
)

type Account struct {
	Id               uint64    `orm:"auto" valid:"Required"`
	AccessKey        string    `orm:"size(20)" valid:"Required"`
	SecretKey        string    `orm:"size(20)" valid:"Required"`
	OrganizationName string    `orm:"size(50)" valid:"Required"`
	UpdateTime       time.Time `orm:"auto_now;type(datetime)"`
	CreateTime       time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *Account) TableUnique() [][]string {
	return [][]string{
		[]string{"AccessKey"},
	}
}

//注册用户
func InsertAccount(account Account) {
	db := GetOrmer()
	db.Insert(&account)
}
