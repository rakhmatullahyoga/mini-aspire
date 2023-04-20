package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/auth"
	mocks "github.com/rakhmatullahyoga/mini-aspire/mocks/auth"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	uc      *mocks.AuthUsecase
	handler http.Handler
}

func (s *AuthHandlerTestSuite) SetupTest() {
	s.uc = new(mocks.AuthUsecase)
	s.handler = auth.NewHandler(s.uc).Router()
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (s *AuthHandlerTestSuite) TestLoginInvalidParameter() {
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.ErrorResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid parameter", resp.Message)
}

func (s *AuthHandlerTestSuite) TestLoginFailed() {
	username := "user"
	password := "pass"
	body := map[string]interface{}{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	s.uc.On("Login", username, password).Return(auth.Token(""), errors.New("invalid username or password")).Once()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.ErrorResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusUnauthorized, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid username or password", resp.Message)
}

func (s *AuthHandlerTestSuite) TestLoginSuccess() {
	username := "user"
	password := "pass"
	body := map[string]interface{}{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	s.uc.On("Login", username, password).Return(auth.Token("token"), nil).Once()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.SuccessResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusOK, rr.Code)
	s.Assert().Equal("success", resp.Status)
	s.Assert().Equal(auth.Token("token"), resp.Token)
}

func (s *AuthHandlerTestSuite) TestRegisterInvalidParameter() {
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.ErrorResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid parameter", resp.Message)
}

func (s *AuthHandlerTestSuite) TestRegisterFailed() {
	username := "user"
	password := "pass"
	body := map[string]interface{}{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	s.uc.On("Register", username, password).Return(auth.Token(""), errors.New("user already existed")).Once()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.ErrorResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusUnprocessableEntity, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("user already existed", resp.Message)
}

func (s *AuthHandlerTestSuite) TestRegisterSuccess() {
	username := "user"
	password := "pass"
	body := map[string]interface{}{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	s.uc.On("Register", username, password).Return(auth.Token("token"), nil).Once()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp auth.SuccessResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	s.Assert().Equal(http.StatusOK, rr.Code)
	s.Assert().Equal("success", resp.Status)
	s.Assert().Equal(auth.Token("token"), resp.Token)
}
