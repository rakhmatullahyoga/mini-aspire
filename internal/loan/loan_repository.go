package loan

import (
	"strconv"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

type LoanRepository struct {
	loans []Loan
}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{
		loans: []Loan{},
	}
}

func (r *LoanRepository) StoreLoan(loan *Loan) error {
	loan.ID = strconv.Itoa(len(r.loans))
	r.loans = append(r.loans, *loan)
	return nil
}

func (r *LoanRepository) ListUserLoans(userID string, offset, count int) ([]Loan, error) {
	var offsetCounter int
	userLoans := []Loan{}
	for i := 0; i < len(r.loans) && len(userLoans) < count; i++ {
		loan := r.loans[i]
		if loan.UserID == userID {
			if offsetCounter >= offset {
				userLoans = append(userLoans, loan)
			}
			offsetCounter++
		}
	}
	return userLoans, nil
}

func (r *LoanRepository) GetLoanByID(loanID string) (*Loan, error) {
	for _, loan := range r.loans {
		if loan.ID == loanID {
			return &loan, nil
		}
	}
	return nil, commons.ErrRecordNotFound
}

func (r *LoanRepository) ListLoans(offset, count int) ([]Loan, error) {
	var offsetCounter int
	loans := []Loan{}
	for i := 0; i < len(r.loans) && len(loans) < count; i++ {
		loan := r.loans[i]
		if offsetCounter >= offset {
			loans = append(loans, loan)
		}
		offsetCounter++
	}
	return loans, nil
}

func (r *LoanRepository) SetLoanStatus(loanID string, status LoanStatus) (*Loan, error) {
	for i, loan := range r.loans {
		if loan.ID == loanID {
			loan.Status = status
			r.loans[i] = loan
			return &loan, nil
		}
	}
	return nil, commons.ErrRecordNotFound
}
