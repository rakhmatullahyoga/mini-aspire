package loan

import (
	"strconv"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

type RepaymentRepository struct {
	repayments []Repayment
}

func NewRepaymentRepository() *RepaymentRepository {
	return &RepaymentRepository{
		repayments: []Repayment{},
	}
}

func (r *RepaymentRepository) BulkStoreRepayment(repayments []Repayment) error {
	for _, repayment := range repayments {
		repayment.ID = strconv.Itoa(len(r.repayments))
		r.repayments = append(r.repayments, repayment)
	}
	return nil
}

func (r *RepaymentRepository) ListLoanRepayments(loanID string, offset, count int) ([]Repayment, error) {
	var offsetCounter int
	repayments := []Repayment{}
	for i := 0; i < len(r.repayments) && len(repayments) < count; i++ {
		repayment := r.repayments[i]
		if repayment.LoanID == loanID {
			if offsetCounter >= offset {
				repayments = append(repayments, repayment)
			}
			offsetCounter++
		}
	}
	return repayments, nil
}

func (r *RepaymentRepository) GetFirstPendingRepayment(loanID string) (*Repayment, error) {
	for _, repayment := range r.repayments {
		if repayment.LoanID == loanID && repayment.Status == RepaymentStatusPending {
			return &repayment, nil
		}
	}
	return nil, commons.ErrRecordNotFound
}

func (r *RepaymentRepository) UpdateRepayment(repayment Repayment) error {
	idx, err := strconv.Atoi(repayment.ID)
	if err != nil || idx >= len(r.repayments) {
		return commons.ErrInvalidRecord
	}
	r.repayments[idx] = repayment
	return nil
}
