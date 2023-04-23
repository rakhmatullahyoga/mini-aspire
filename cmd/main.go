package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rakhmatullahyoga/mini-aspire/internal/auth"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
)

func main() {
	userRepo := auth.NewUserRepository()
	authUsecase := auth.NewUsecase(userRepo)
	authHandler := auth.NewHandler(authUsecase)

	loanRepo := loan.NewLoanRepository()
	repaymentRepo := loan.NewRepaymentRepository()
	loanUsecase := loan.NewUsecase(loanRepo, repaymentRepo)
	loanHandler := loan.NewUserHandler(loanUsecase)
	adminLoanHandler := loan.NewAdminHandler(loanUsecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Mount("/auth", authHandler.Router())
	r.Mount("/loans", loanHandler.Router())
	r.Mount("/admin", adminLoanHandler.Router())
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalf("cannot serve from localhost:8080 : %s\n", err.Error())
	}
}
