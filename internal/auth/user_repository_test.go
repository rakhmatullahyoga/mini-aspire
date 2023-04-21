package auth_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTestSuite struct {
	suite.Suite
	repo *auth.UserRepository
}

func (s *userRepositoryTestSuite) SetupTest() {
	s.repo = auth.NewUserRepository()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}

func (s *userRepositoryTestSuite) TestFindByUsernameNotFound() {
	user, err := s.repo.FindByUsername("nonadmin")
	s.Assert().NotNil(err)
	s.Assert().Equal(commons.ErrRecordNotFound.Error(), err.Error())
	s.Assert().Nil(user)
}

func (s *userRepositoryTestSuite) TestFindByUsernameFound() {
	user, err := s.repo.FindByUsername("admin")
	s.Assert().Nil(err)
	s.Assert().NotNil(user)
	s.Assert().Equal(auth.Username("admin"), user.Username)
}

func (s *userRepositoryTestSuite) TestStoreUser() {
	user := auth.User{
		Username: auth.Username("user1"),
		Password: auth.Password("pass"),
	}
	err := s.repo.StoreUser(user)
	userStored, _ := s.repo.FindByUsername("user1")
	s.Assert().Nil(err)
	s.Assert().NotNil(userStored)
	s.Assert().Equal(user.Username, userStored.Username)
	s.Assert().Equal(user.Password, userStored.Password)
	s.Assert().Equal(false, userStored.IsAdmin)
}
