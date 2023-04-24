package loan_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/stretchr/testify/suite"
)

type loanRepositoryTestSuite struct {
	suite.Suite
	repo *loan.LoanRepository
}

func (s *loanRepositoryTestSuite) SetupTest() {
	s.repo = loan.NewLoanRepository()
}

func TestLoanRepository(t *testing.T) {
	suite.Run(t, new(loanRepositoryTestSuite))
}

func (s *loanRepositoryTestSuite) TestStoreLoan() {
	newLoan := loan.Loan{
		UserID: "123",
		Amount: 5000,
		Term:   12,
	}
	err := s.repo.StoreLoan(&newLoan)
	s.Assert().Nil(err)
}

func (s *loanRepositoryTestSuite) TestListUserLoansNoData() {
	loanList, err := s.repo.ListLoans(0, 10)
	s.Assert().Nil(err)
	s.Assert().Empty(loanList)
}

func (s *loanRepositoryTestSuite) TestListUserLoansFound() {
	newLoan := loan.Loan{
		UserID: "123",
		Amount: 5000,
		Term:   12,
	}
	s.repo.StoreLoan(&newLoan)
	loanList, err := s.repo.ListLoans(0, 10)
	s.Assert().Nil(err)
	s.Assert().NotEmpty(loanList)
	s.Assert().Equal(newLoan, loanList[0])
}

func (s *loanRepositoryTestSuite) TestGetLoanByIDNotFound() {
	loan, err := s.repo.GetLoanByID("1")
	s.Assert().NotNil(err)
	s.Assert().Equal("record not found", err.Error())
	s.Assert().Nil(loan)
}

func (s *loanRepositoryTestSuite) TestGetLoanByIDFound() {
	newLoan := loan.Loan{
		UserID: "1",
		Amount: 5000,
		Term:   12,
	}
	s.repo.StoreLoan(&newLoan)
	loan, err := s.repo.GetLoanByID(newLoan.ID)
	s.Assert().Nil(err)
	s.Assert().NotNil(loan)
	s.Assert().Equal(newLoan, *loan)
}

func (s *loanRepositoryTestSuite) TestListLoansNoData() {}

func (s *loanRepositoryTestSuite) TestListLoansFound() {}

func (s *loanRepositoryTestSuite) TestSetLoanStatusNotFound() {}

func (s *loanRepositoryTestSuite) TestSetLoanStatusSuccess() {}
