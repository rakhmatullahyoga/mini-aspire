package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Username string
type Password string
type Token string

type User struct {
	Username Username
	Password Password
	IsAdmin  bool
}

type Claims struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

var (
	errInvalidLogin = errors.New("invalid username or password")
	jwtKey          = []byte("some_secret_key")
)

type UserRepository interface {
	FindByUsername(username string) *User
	StoreUser(user User) error
}

type authUsecase struct {
	repo UserRepository
}

func NewUsecase(repo UserRepository) *authUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (u *authUsecase) Login(username, password string) (Token, error) {
	user := u.repo.FindByUsername(username)
	if user == nil || user.Password != Password(password) {
		return "", errInvalidLogin
	}

	return generateToken(*user)
}

func (u *authUsecase) Register(username, password string) (Token, error) {
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
	claims := &Claims{
		UserID:           string(user.Username),
		IsAdmin:          user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	return Token(tokenStr), err
}
