package models

import "time"

type Comment struct {
	Id            int
	UserId        int
	NickName      string
	ReplyUserId   int
	ReplyNickName string
	ArticleId     int
	Content       string
	CreatedAt     time.Time `gorm:"column:createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"`
	State         int8
}

func (Comment) TableName() string {
	return "cms_comment"
}
