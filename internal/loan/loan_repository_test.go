package loan_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/stretchr/testify/assert"
)

func TestLoanRepository_StoreLoan(t *testing.T) {
	repo := loan.NewLoanRepository()

	// create new loan
	newLoan := loan.Loan{
		UserID: "123",
		Amount: 5000,
		Term:   12,
	}
	err := repo.StoreLoan(&newLoan)

	// assert no error
	assert.NoError(t, err)

	// assert loan is stored in repository
	loanList, _ := repo.ListLoans(0, 10)
	assert.Equal(t, []loan.Loan{newLoan}, loanList)

	// create another loan
	anotherLoan := loan.Loan{
		UserID: "123",
		Amount: 3000,
		Term:   6,
	}
	err = repo.StoreLoan(&anotherLoan)

	// assert no error
	assert.NoError(t, err)

	// assert both loans are stored in repository
	loanList, _ = repo.ListLoans(0, 10)
	assert.Equal(t, []loan.Loan{newLoan, anotherLoan}, loanList)
}
