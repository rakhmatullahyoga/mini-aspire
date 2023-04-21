package auth_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	username = "user"
	password = "pass"
)

type authUsecaseTestSuite struct {
	suite.Suite
	repo *mocks.IUserRepository
	uc   *auth.AuthUsecase
}

func (s *authUsecaseTestSuite) SetupTest() {
	s.repo = new(mocks.IUserRepository)
	s.uc = auth.NewUsecase(s.repo)
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(authUsecaseTestSuite))
}

func (s *authUsecaseTestSuite) TestLoginUserNotFound() {
	s.repo.On("FindByUsername", username).Return(nil, errors.New("user not found")).Once()
	_, err := s.uc.Login(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("invalid username or password", err.Error())
}

func (s *authUsecaseTestSuite) TestLoginWrongPassword() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password("password"),
	}
	s.repo.On("FindByUsername", username).Return(&user, nil).Once()
	_, err := s.uc.Login(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("invalid username or password", err.Error())
}

func (s *authUsecaseTestSuite) TestLoginSuccess() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password(password),
	}
	s.repo.On("FindByUsername", username).Return(&user, nil).Once()
	token, err := s.uc.Login(username, password)
	s.Assert().Nil(err)
	s.Assert().NotNil(token)
}

func (s *authUsecaseTestSuite) TestRegisterExistingUser() {
	user := auth.User{
		Username: auth.Username(username),
		Password: auth.Password(password),
	}
	s.repo.On("FindByUsername", username).Return(&user, nil).Once()
	_, err := s.uc.Register(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal(fmt.Sprintf("username %s already registered", username), err.Error())
}

func (s *authUsecaseTestSuite) TestRegisterError() {
	s.repo.On("FindByUsername", username).Return(nil, nil).Once()
	s.repo.On("StoreUser", mock.Anything).Return(errors.New("unexpected")).Once()
	_, err := s.uc.Register(username, password)
	s.Assert().NotNil(err)
	s.Assert().Equal("unexpected", err.Error())
}

func (s *authUsecaseTestSuite) TestRegisterSuccess() {
	s.repo.On("FindByUsername", username).Return(nil, nil).Once()
	s.repo.On("StoreUser", mock.Anything).Return(nil).Once()
	token, err := s.uc.Register(username, password)
	s.Assert().Nil(err)
	s.Assert().NotNil(token)
}
