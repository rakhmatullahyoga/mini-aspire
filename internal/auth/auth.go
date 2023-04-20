package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

type Username string
type Password string
type Token string

type User struct {
	Username Username
	Password Password
	IsAdmin  bool
}

var (
	errInvalidLogin = errors.New("invalid username or password")
)

//go:generate mockery --name=IUserRepository --structname IUserRepository --filename=IUserRepository.go --output=mocks
type IUserRepository interface {
	FindByUsername(username string) *User
	StoreUser(user User) error
}

type AuthUsecase struct {
	repo IUserRepository
}

func NewUsecase(repo IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (u *AuthUsecase) Login(username, password string) (Token, error) {
	user := u.repo.FindByUsername(username)
	if user == nil || user.Password != Password(password) {
		return "", errInvalidLogin
	}

	return generateToken(*user)
}

func (u *AuthUsecase) Register(username, password string) (Token, error) {
	existUser := u.repo.FindByUsername(username)
	if existUser != nil {
		return "", fmt.Errorf("username %s already registered", username)
	}

	user := User{
		Username: Username(username),
		Password: Password(password),
		IsAdmin:  false,
	}
	err := u.repo.StoreUser(user)
	if err != nil {
		return "", err
	}

	return generateToken(user)
}

func generateToken(user User) (Token, error) {
	claims := &commons.Claims{
		UserID:           string(user.Username),
		IsAdmin:          user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(commons.JwtKey)
	return Token(tokenStr), err
}