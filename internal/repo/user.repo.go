package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUser() string {
	// Get user ID from the request parameters
	return "Get user Hung Nguyen"
}
