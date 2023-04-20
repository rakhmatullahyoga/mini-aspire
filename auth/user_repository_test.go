package auth_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/auth"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repo *auth.UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	s.repo = auth.NewUserRepository()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestFindByUsernameNotFound() {
	user := s.repo.FindByUsername("nonadmin")
	s.Assert().Nil(user)
}

func (s *UserRepositoryTestSuite) TestFindByUsernameFound() {
	user := s.repo.FindByUsername("admin")
	s.Assert().NotNil(user)
	s.Assert().Equal(auth.Username("admin"), user.Username)
}

func (s *UserRepositoryTestSuite) TestStoreUser() {
	user := auth.User{
		Username: auth.Username("user1"),
		Password: auth.Password("pass"),
	}
	err := s.repo.StoreUser(user)
	userStored := s.repo.FindByUsername("user1")
	s.Assert().Nil(err)
	s.Assert().NotNil(userStored)
	s.Assert().Equal(user.Username, userStored.Username)
	s.Assert().Equal(user.Password, userStored.Password)
	s.Assert().Equal(false, userStored.IsAdmin)
}
