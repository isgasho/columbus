package model

import (
	//"errors"
	//"fmt"

	//"github.com/satori/go.uuid"
)

type Account struct {
	Id   int       `json:"id" xorm:"pk autoincr"`
	Name string    `json:"name" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
}


func (this *Account) TableName() string {
	return "account"
}

