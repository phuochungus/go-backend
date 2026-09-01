package service

import "go-backend/internal/repo"

// type UserService struct {
// 	repo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		repo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUser() string {
// 	// Get user ID from the request parameters
// 	return us.repo.GetUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	repo repo.IUserRepository
}

func NewUserService(repo repo.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(email string, purpose string) int {
	panic("umimplement")
}
