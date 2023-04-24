package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
)

func TestValidateJWT(t *testing.T) {
	// Create a fake handler to be passed to validateJWT
	fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Test case: Authorization header is missing, should return 401 Unauthorized
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := auth.ValidateJWT(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("validateJWT returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	// Test case: JWT is invalid, should return 401 Unauthorized
	req, err = http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer invalid_token")
	rr = httptest.NewRecorder()
	handler = auth.ValidateJWT(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("validateJWT returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Test case: JWT signature is invalid, should return 401 Unauthorized
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(commons.ClaimsKeyUserID):  "123",
		string(commons.ClaimsKeyIsAdmin): true,
	}).SignedString([]byte("invalid_key"))
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	rr = httptest.NewRecorder()
	handler = auth.ValidateJWT(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("validateJWT returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Test case: JWT is valid, should pass the request to the next handler
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(commons.ClaimsKeyUserID):  "123",
		string(commons.ClaimsKeyIsAdmin): true,
	}).SignedString([]byte(commons.JwtKey))
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	rr = httptest.NewRecorder()
	handler = auth.ValidateJWT(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("validateJWT returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestEnsureAdmin(t *testing.T) {
	// Create a fake handler to be passed to ensureAdmin
	fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Test case: User is not an admin, should return 403 Forbidden
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(req.Context(), commons.ClaimsKeyIsAdmin, false)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	handler := auth.EnsureAdmin(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("ensureAdmin returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}

	// Test case: User is an admin, should pass the request to the next handler
	req, err = http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx = context.WithValue(req.Context(), commons.ClaimsKeyIsAdmin, true)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler = auth.EnsureAdmin(fakeHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ensureAdmin returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
