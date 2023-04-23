package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth/mocks"
	"github.com/stretchr/testify/suite"
)

type authHandlerTestSuite struct {
	suite.Suite
	uc      *mocks.IAuthUsecase
	handler http.Handler
}

func (s *authHandlerTestSuite) SetupTest() {
	s.uc = new(mocks.IAuthUsecase)
	s.handler = auth.NewHandler(s.uc).Router()
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(authHandlerTestSuite))
}

func (s *authHandlerTestSuite) TestLoginInvalidRequestBody() {
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid request body", resp.Message)
}

func (s *authHandlerTestSuite) TestLoginInvalidParameter() {
	username := "user"
	body := map[string]interface{}{
		"username": username,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid parameter", resp.Message)
}

func (s *authHandlerTestSuite) TestLoginFailed() {
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
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusUnauthorized, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid username or password", resp.Message)
}

func (s *authHandlerTestSuite) TestLoginSuccess() {
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
	var resp commons.SuccessResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"token": "token",
	}
	s.Assert().Equal(http.StatusOK, rr.Code)
	s.Assert().Equal("success", resp.Status)
	s.Assert().Equal(data, resp.Data)
}

func (s *authHandlerTestSuite) TestRegisterInvalidRequestBody() {
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid request body", resp.Message)
}

func (s *authHandlerTestSuite) TestRegisterInvalidParameter() {
	username := "user"
	body := map[string]interface{}{
		"username": username,
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusBadRequest, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("invalid parameter", resp.Message)
}

func (s *authHandlerTestSuite) TestRegisterFailed() {
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
	var resp commons.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	s.Assert().Equal(http.StatusUnprocessableEntity, rr.Code)
	s.Assert().Equal("failed", resp.Status)
	s.Assert().Equal("user already existed", resp.Message)
}

func (s *authHandlerTestSuite) TestRegisterSuccess() {
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
	var resp commons.SuccessResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"token": "token",
	}
	s.Assert().Equal(http.StatusOK, rr.Code)
	s.Assert().Equal("success", resp.Status)
	s.Assert().Equal(data, resp.Data)
}
