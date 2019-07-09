package models

import (
	"time"
)

type Article struct {
	Id        int
	Title     string
	Source    string
	Summary   string
	Keyword   string
	PicUrl    string
	Content   string
	Status    int8
	AuditId   int
	Auditor   string
	AuditAt   time.Time
	AuditNote string
	Stars     int
	Comments  int
	Sn        int
	CreateId  int `gorm:"column:createId"`
	Creator   string
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
	State     int8
}

func (Article) TableName() string {
	return "cms_article"
}
