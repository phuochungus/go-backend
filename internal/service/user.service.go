package service

import "go-backend/internal/repo"

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUser() string {
	// Get user ID from the request parameters
	return us.repo.GetUser()
}
