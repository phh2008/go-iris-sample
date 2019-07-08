package models

import (
	"time"
)

type User struct {
	Id        int
	UserName  string
	Password  string
	NickName  string
	RealName  string
	Mobile    string
	Email     string
	Sex       int8
	Type      int8
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	State     int8
}

func (User) TableName() string {
	return "mem_user"
}
