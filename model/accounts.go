package model

import (
	//"errors"
	//"fmt"
	"time"
	//"github.com/satori/go.uuid"
)

// 用户基本信息
type Account struct {
	Uid         int       `json:"uid" xorm:"pk autoincr"`
	Name        string    `json:"name" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	Email       string    `json:"email"`
	Avatar      string    `json:"avatar"`
	City        string    `json:"city"`
	Introduce   string    `json:"introduce"`
	Status      int       `json:"status"`
	Ctime       time.Time `json:"ctime" xorm:"created"`
	Mtime       time.Time `json:"mtime" xorm:"<-"`
}

const (
	Approach = "未删"
	FdelHasDel = "已删"
)

// 授权凭证
type Authentication struct {
	Id          int       `json:"id" xorm:"pk autoincr"`
	Uid         int       `json:"uid"`
    Approach    string    ``
	Ctime       time.Time `json:"ctime" xorm:"created"`
	Mtime       time.Time `json:"mtime" xorm:"<-"`
}

func (this *Account) TableName() string {
	return "account"
}

