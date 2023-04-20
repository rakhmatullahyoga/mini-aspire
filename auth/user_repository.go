package auth

type userRepository struct {
	users map[Username]User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: map[Username]User{
			Username("admin"): {
				Username: "admin",
				Password: "admin",
				IsAdmin:  true,
			},
		},
	}
}

func (r *userRepository) FindByUsername(username string) *User {
	user, ok := r.users[Username(username)]
	if !ok {
		return nil
	}

	return &user
}

func (r *userRepository) StoreUser(user User) error {
	r.users[user.Username] = user
	return nil
}
