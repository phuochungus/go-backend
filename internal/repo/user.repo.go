package repo

import "go-backend/internal/po"

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUser() string {
// 	// Get user ID from the request parameters
// 	return "Get user Hung Nguyen"
// }

type IUserRepository interface {
	GetUserByEmail(email string) (*po.User, error)
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByEmail(email string) (*po.User, error) {
	panic("unimplemented")
}
