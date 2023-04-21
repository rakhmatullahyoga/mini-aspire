package loan

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

type LoanRequest struct {
	Amount   int       `json:"amount" validate:"required"`
	Term     int       `json:"term"`
	LoanDate time.Time `json:"loan_date"`
}

//go:generate mockery --name=ILoanUsecase --structname ILoanUsecase --filename=ILoanUsecase.go --output=mocks
type ILoanUsecase interface {
	SubmitLoan(ctx context.Context, req LoanRequest) (*Loan, error)
	ListUserLoans(ctx context.Context, page int) ([]Loan, error)
	UserGetLoan(ctx context.Context, loanID string) (*Loan, error)
	ListUserRepayments(ctx context.Context, loanID string, page int) ([]Repayment, error)
	UserSubmitRepayment(ctx context.Context, loanID string, amount float64) (*Repayment, error)
}

type userHandler struct {
	uc ILoanUsecase
}

func NewUserHandler(uc ILoanUsecase) *userHandler {
	return &userHandler{
		uc: uc,
	}
}

func (h *userHandler) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(validateJWT)
	r.Post("/", h.submitLoan)
	r.Get("/", h.listUserLoans)
	r.Get("/{loanID}", h.userGetLoan)
	r.Get("/{loanID}/repayments", h.listUserRepayments)
	r.Post("/{loanID}/repayments", h.userSubmitRepayment)
	return r
}

func (h *userHandler) submitLoan(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	loanReq := LoanRequest{
		Term:     1,
		LoanDate: today,
	}
	err := json.NewDecoder(r.Body).Decode(&loanReq)
	if err != nil {
		res := commons.BuildErrorResponse("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	validate := validator.New()
	err = validate.Struct(loanReq)
	if err != nil {
		res := commons.BuildErrorResponse("invalid parameter")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	loan, err := h.uc.SubmitLoan(r.Context(), loanReq)
	if err != nil {
		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(loan)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *userHandler) listUserLoans(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	pageStr := val.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	loans, err := h.uc.ListUserLoans(r.Context(), page)
	if err != nil {
		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(loans)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *userHandler) userGetLoan(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanID")
	loan, err := h.uc.UserGetLoan(r.Context(), loanID)
	if err != nil {
		var httpStatus int
		switch err {
		case errLoanOwnership:
			httpStatus = http.StatusForbidden
		case commons.ErrRecordNotFound:
			httpStatus = http.StatusNotFound
		default:
			httpStatus = http.StatusInternalServerError
		}

		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(httpStatus)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(loan)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *userHandler) listUserRepayments(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanID")
	val := r.URL.Query()
	pageStr := val.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	repayments, err := h.uc.ListUserRepayments(r.Context(), loanID, page)
	if err != nil {
		var httpStatus int
		switch err {
		case errLoanOwnership:
			httpStatus = http.StatusForbidden
		case commons.ErrRecordNotFound:
			httpStatus = http.StatusNotFound
		default:
			httpStatus = http.StatusInternalServerError
		}

		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(httpStatus)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(repayments)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *userHandler) userSubmitRepayment(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanID")
	var reqBody struct {
		Amount float64 `json:"amount" validate:"required"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		res := commons.BuildErrorResponse("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	validate := validator.New()
	err = validate.Struct(reqBody)
	if err != nil {
		res := commons.BuildErrorResponse("invalid parameter")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	repayment, err := h.uc.UserSubmitRepayment(r.Context(), loanID, reqBody.Amount)
	if err != nil {
		var httpStatus int
		switch err {
		case errLoanOwnership:
			httpStatus = http.StatusForbidden
		case errLoanNotApproved, errInsufficient, errLoanAlreadyPaid:
			httpStatus = http.StatusUnprocessableEntity
		case commons.ErrRecordNotFound:
			httpStatus = http.StatusNotFound
		default:
			httpStatus = http.StatusInternalServerError
		}
		res := commons.BuildErrorResponse(err.Error())
		w.WriteHeader(httpStatus)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(repayment)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
