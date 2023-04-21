package auth

import "github.com/rakhmatullahyoga/mini-aspire/commons"

type UserRepository struct {
	users map[Username]User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[Username]User{
			Username("admin"): {
				Username: "admin",
				Password: "admin",
				IsAdmin:  true,
			},
		},
	}
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	user, ok := r.users[Username(username)]
	if !ok {
		return nil, commons.ErrRecordNotFound
	}

	return &user, nil
}

func (r *UserRepository) StoreUser(user User) error {
	r.users[user.Username] = user
	return nil
}
