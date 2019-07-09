package models

import "time"

type Star struct {
	Id        int
	UserId    int
	ArticleId int
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
	State     int8
}

func (Star) TableName() string {
	return "cms_star"
}
