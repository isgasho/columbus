package models

import (
	"time"
)

type Account struct {
	Uid       int       `json:"uid" xorm:"not null pk autoincr INT(10)"`
	Email     string    `json:"email" xorm:"not null default '' unique VARCHAR(128)"`
	Username  string    `json:"username" xorm:"not null comment('用户名') unique VARCHAR(20)"`
	Name      string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(20)"`
	Mobile    string    `json:"mobile" xorm:"not null default '' comment('手机号') VARCHAR(20)"`
	Avatar    string    `json:"avatar" xorm:"not null default '' comment('头像(如果为空，则使用http://www.gravatar.com)') VARCHAR(128)"`
	City      string    `json:"city" xorm:"not null default '' comment('居住地') VARCHAR(10)"`
	Website   string    `json:"website" xorm:"not null default '' comment('个人主页，博客') VARCHAR(63)"`
	Monlog    string    `json:"monlog" xorm:"not null default '' comment('个人状态，签名，独白') VARCHAR(140)"`
	Introduce string    `json:"introduce" xorm:"not null comment('个人简介') VARCHAR(2022)"`
	IsRoot    int       `json:"is_root" xorm:"not null default 0 comment('是否超级用户，不受权限控制：1-是') TINYINT(3)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func (this *Account) TableName() string {
	return "account"
}