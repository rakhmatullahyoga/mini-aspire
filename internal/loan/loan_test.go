package loan_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan/mocks"
	"github.com/stretchr/testify/suite"
)

type loanUsecaseTestSuite struct {
	suite.Suite
	loanRepo      *mocks.ILoanRepository
	repaymentRepo *mocks.IRepaymentRepository
	uc            *loan.LoanUsecase
}

func (s *loanUsecaseTestSuite) SetupTest() {
	s.loanRepo = new(mocks.ILoanRepository)
	s.repaymentRepo = new(mocks.IRepaymentRepository)
	s.uc = loan.NewUsecase(s.loanRepo, s.repaymentRepo)
}

func TestLoanUsecase(t *testing.T) {
	suite.Run(t, new(loanUsecaseTestSuite))
}
