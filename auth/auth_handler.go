package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type credentials struct {
	Password *string `json:"password"`
	Username *string `json:"username"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Token  Token  `json:"token"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AuthUsecase interface {
	Login(username, password string) (Token, error)
	Register(username, password string) (Token, error)
}

type handler struct {
	uc AuthUsecase
}

func NewHandler(uc AuthUsecase) *handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) AuthRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", h.login)
	r.Post("/register", h.register)
	return r
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == nil || creds.Password == nil {
		res := ErrorResponse{
			Status:  "failed",
			Message: "invalid parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	token, err := h.uc.Login(*creds.Username, *creds.Password)
	if err != nil {
		res := ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := SuccessResponse{
		Status: "success",
		Token:  token,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == nil || creds.Password == nil {
		res := ErrorResponse{
			Status:  "failed",
			Message: "invalid parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	token, err := h.uc.Register(*creds.Username, *creds.Password)
	if err != nil {
		res := ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := SuccessResponse{
		Status: "success",
		Token:  token,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
