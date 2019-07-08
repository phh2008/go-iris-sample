package services

import "com.phh/blog/repositories"

type UserService interface {
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepo
}
