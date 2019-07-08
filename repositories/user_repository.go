package repositories

import "github.com/jinzhu/gorm"

type UserRepo interface {
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}
