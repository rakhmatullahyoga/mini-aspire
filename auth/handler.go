package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

type credentials struct {
	Password *string `json:"password" validate:"required"`
	Username *string `json:"username" validate:"required"`
}

//go:generate mockery --name=AuthUsecase --structname AuthUsecase --filename=AuthUsecase.go --output=mocks
type IAuthUsecase interface {
	Login(username, password string) (Token, error)
	Register(username, password string) (Token, error)
}

type handler struct {
	uc IAuthUsecase
}

func NewHandler(uc IAuthUsecase) *handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", h.login)
	r.Post("/register", h.register)
	return r
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		res := commons.BuildErrorResponse("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	validate := validator.New()
	err = validate.Struct(creds)
	if err != nil {
		res := commons.BuildErrorResponse("invalid parameter")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	token, err := h.uc.Login(*creds.Username, *creds.Password)
	if err != nil {
		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.SuccessResponse{
		Status: "success",
		Data: struct {
			Token Token `json:"token"`
		}{
			Token: token,
		},
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		res := commons.BuildErrorResponse("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	validate := validator.New()
	err = validate.Struct(creds)
	if err != nil {
		res := commons.BuildErrorResponse("invalid parameter")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	token, err := h.uc.Register(*creds.Username, *creds.Password)
	if err != nil {
		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.SuccessResponse{
		Status: "success",
		Data: struct {
			Token Token `json:"token"`
		}{
			Token: token,
		},
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
