package loan

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

//go:generate mockery --name=IAdminUsecase --structname IAdminUsecase --filename=IAdminUsecase.go --output=mocks
type IAdminUsecase interface {
	ListLoans(page int) ([]Loan, error)
	ApproveLoan(loanID string) (*Loan, error)
}

type adminHandler struct {
	uc IAdminUsecase
}

func NewAdminHandler(uc IAdminUsecase) *adminHandler {
	return &adminHandler{
		uc: uc,
	}
}

func (h *adminHandler) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(validateJWT)
	r.Use(ensureAdmin)
	r.Get("/loans", h.adminListLoans)
	r.Post("/loans/{loanID}", h.adminApproveLoan)
	return r
}

func (h *adminHandler) adminListLoans(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	pageStr := val.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	loans, err := h.uc.ListLoans(page)
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

func (h *adminHandler) adminApproveLoan(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "loanID")
	loan, err := h.uc.ApproveLoan(loanID)
	if err != nil {
		var httpStatus int
		res := commons.BuildErrorResponse(err.Error())
		if err == commons.ErrRecordNotFound {
			httpStatus = http.StatusNotFound
		} else {
			httpStatus = http.StatusInternalServerError
		}
		w.WriteHeader(httpStatus)
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res := commons.BuildResponse(loan)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
