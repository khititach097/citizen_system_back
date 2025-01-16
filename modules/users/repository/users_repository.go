package user_repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FetchAllUsers() (string, error) {
	// Example: Return mock data for simplicity
	return `["User1", "User2"]`, nil
}
