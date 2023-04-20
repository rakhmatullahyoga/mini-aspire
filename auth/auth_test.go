package auth_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/auth"
	"github.com/rakhmatullahyoga/mini-aspire/auth/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	username = "user"
	password = "pass"
)

type AuthUsecaseTestSuite struct {
	suite.Suite
	repo *mocks.UserRepository
	uc   auth.IAuthUsecase
}

func (s *AuthUsecaseTestSuite) SetupTest() {
	s.repo = new(mocks.UserRepository)
	s.uc = auth.NewUsecase(s.repo)
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}

func (s *AuthUsecaseTestSuite) TestLoginUserNotFound() {
	s.repo.On("FindByUsername", username).Return(nil).Once()
	_, err := s.uc.Login(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("invalid username or password", err.Error())
}

func (s *AuthUsecaseTestSuite) TestLoginWrongPassword() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password("password"),
	}
	s.repo.On("FindByUsername", username).Return(&user).Once()
	_, err := s.uc.Login(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("invalid username or password", err.Error())
}

func (s *AuthUsecaseTestSuite) TestLoginSuccess() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password(password),
	}
	s.repo.On("FindByUsername", username).Return(&user).Once()
	token, err := s.uc.Login(username, password)
	s.Assert().Nil(err)
	s.Assert().NotNil(token)
}

func (s *AuthUsecaseTestSuite) TestRegisterExistingUser() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password(password),
	}
	s.repo.On("FindByUsername", username).Return(&user).Once()
	_, err := s.uc.Register(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal(fmt.Sprintf("username %s already registered", username), err.Error())
}

func (s *AuthUsecaseTestSuite) TestRegisterError() {
	s.repo.On("FindByUsername", username).Return(nil).Once()
	s.repo.On("StoreUser", mock.Anything).Return(errors.New("unexpected")).Once()
	_, err := s.uc.Register(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("unexpected", err.Error())
}

func (s *AuthUsecaseTestSuite) TestRegisterSuccess() {
	s.repo.On("FindByUsername", username).Return(nil).Once()
	s.repo.On("StoreUser", mock.Anything).Return(nil).Once()
	token, err := s.uc.Register(username, password)
	s.Assert().Nil(err)
	s.Assert().NotNil(token)
}
