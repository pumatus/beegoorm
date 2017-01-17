package models

import (
	"time"
)

type RoleStatus uint8

type Role struct {
	Id         uint64 `orm:"auto"`
	Name       string `orm:"size(20)"`
	Status     RoleStatus
	UpdateTime time.Time `orm:"auto_now";"type(datetime)"`
	CreateTime time.Time `orm:"auto_now_add";"type(datetime)"`
}

const (
	UNACTIVE RoleStatus = iota
	ACTIVE
)
