package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
)

func main() {
	repoUser := auth.NewUserRepository()
	ucAuth := auth.NewUsecase(repoUser)
	handlerAuth := auth.NewHandler(ucAuth)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Mount("/auth", handlerAuth.Router())
	http.ListenAndServe("localhost:8080", r)
}
