package models

import (
	"time"
)

type Authentication struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	OpenId     string    `json:"open_id" xorm:"not null default '' comment('用户的标识，对当前公众号/小程序唯一') unique VARCHAR(127)"`
	UnionId    string    `json:"union_id" xorm:"not null default '' comment('用户的标识，对开放者唯一') VARCHAR(127)"`
	Nickname   string    `json:"nickname" xorm:"not null default '' comment('用户的昵称') VARCHAR(127)"`
	SessionKey string    `json:"session_key" xorm:"not null default '' comment('小程序返回的 session_key') VARCHAR(127)"`
	Avatar     string    `json:"avatar" xorm:"not null default '' comment('用户微信头像') VARCHAR(255)"`
	OpenInfo   string    `json:"open_info" xorm:"not null default '' comment('用户微信的其他信息，json格式') VARCHAR(1024)"`
	Uid        int       `json:"uid" xorm:"not null default 0 comment('用户UID') index INT(11)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' index TIMESTAMP"`
}

func (this *Authentication) TableName() string {
	return "authentication"
}