package loan_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var userId = "user"

type loanUsecaseTestSuite struct {
	suite.Suite
	loanRepo      *mocks.ILoanRepository
	repaymentRepo *mocks.IRepaymentRepository
	uc            *loan.LoanUsecase
	ctx           context.Context
}

func (s *loanUsecaseTestSuite) SetupTest() {
	s.loanRepo = new(mocks.ILoanRepository)
	s.repaymentRepo = new(mocks.IRepaymentRepository)
	s.uc = loan.NewUsecase(s.loanRepo, s.repaymentRepo)
	ctx := context.Background()
	s.ctx = context.WithValue(ctx, commons.ClaimsKeyUserID, userId)
}

func TestLoanUsecase(t *testing.T) {
	suite.Run(t, new(loanUsecaseTestSuite))
}

func (s *loanUsecaseTestSuite) TestSubmitLoanErrorStoreLoan() {
	req := loan.LoanRequest{
		Amount:   10000,
		Term:     3,
		LoanDate: time.Now(),
	}
	s.loanRepo.On("StoreLoan", mock.Anything).Return(errors.New("cannot store loan")).Once()
	loan, err := s.uc.SubmitLoan(s.ctx, req)
	s.Assert().NotNil(err)
	s.Assert().Equal("cannot store loan", err.Error())
	s.Assert().Nil(loan)
}

func (s *loanUsecaseTestSuite) TestSubmitLoanErrorBulkStoreRepayment() {
	req := loan.LoanRequest{
		Amount:   10000,
		Term:     3,
		LoanDate: time.Now(),
	}
	s.loanRepo.On("StoreLoan", mock.Anything).Return(nil).Once()
	s.repaymentRepo.On("BulkStoreRepayment", mock.Anything).Return(errors.New("cannot store repayments")).Once()
	loan, err := s.uc.SubmitLoan(s.ctx, req)
	s.Assert().NotNil(err)
	s.Assert().Equal("cannot store repayments", err.Error())
	s.Assert().Nil(loan)
}

func (s *loanUsecaseTestSuite) TestSubmitLoanSuccess() {
	req := loan.LoanRequest{
		Amount:   10000,
		Term:     3,
		LoanDate: time.Now(),
	}
	s.loanRepo.On("StoreLoan", mock.Anything).Return(nil).Once()
	s.repaymentRepo.On("BulkStoreRepayment", mock.Anything).Return(nil).Once()
	loan, err := s.uc.SubmitLoan(s.ctx, req)
	s.Assert().Nil(err)
	s.Assert().NotNil(loan)
	s.Assert().Equal(10000, loan.Amount)
	s.Assert().Equal(3, loan.Term)
}

func (s *loanUsecaseTestSuite) TestListUserLoans() {
	page := 1
	s.loanRepo.On("ListUserLoans", userId, 0, mock.Anything).Return([]loan.Loan{}, nil).Once()
	loans, err := s.uc.ListUserLoans(s.ctx, page)
	s.Assert().Nil(err)
	s.Assert().Empty(loans)
}

func (s *loanUsecaseTestSuite) TestUserGetLoanErrorRepo() {
	loanID := "1"
	s.loanRepo.On("GetLoanByID", loanID).Return(nil, errors.New("cannot get loan")).Once()
	loan, err := s.uc.UserGetLoan(s.ctx, loanID)
	s.Assert().NotNil(err)
	s.Assert().Equal("cannot get loan", err.Error())
	s.Assert().Nil(loan)
}

func (s *loanUsecaseTestSuite) TestUserGetLoanErrorOwnership() {
	loanID := "1"
	s.loanRepo.On("GetLoanByID", loanID).Return(&loan.Loan{}, nil).Once()
	loan, err := s.uc.UserGetLoan(s.ctx, loanID)
	s.Assert().NotNil(err)
	s.Assert().Equal("invalid resource ownership", err.Error())
	s.Assert().Nil(loan)
}

func (s *loanUsecaseTestSuite) TestUserGetLoanSuccess() {
	loanID := "1"
	loan := &loan.Loan{
		UserID: userId,
	}
	s.loanRepo.On("GetLoanByID", loanID).Return(loan, nil).Once()
	loan, err := s.uc.UserGetLoan(s.ctx, loanID)
	s.Assert().Nil(err)
	s.Assert().NotNil(loan)
}

func (s *loanUsecaseTestSuite) TestListLoans() {
	s.loanRepo.On("ListLoans", mock.Anything, mock.Anything).Return([]loan.Loan{}, nil).Once()
	loans, err := s.uc.ListLoans(1)
	s.Assert().Nil(err)
	s.Assert().Empty(loans)
}

func (s *loanUsecaseTestSuite) TestApproveLoan() {
	loanID := "1"
	s.loanRepo.On("SetLoanStatus", loanID, loan.LoanStatusApproved).Return(&loan.Loan{}, nil).Once()
	loan, err := s.uc.ApproveLoan(loanID)
	s.Assert().Nil(err)
	s.Assert().NotNil(loan)
	s.Assert().Empty(loan)
}

func (s *loanUsecaseTestSuite) TestListUserRepaymentsErrorRepo() {}

func (s *loanUsecaseTestSuite) TestListUserRepaymentsErrorOwnership() {}

func (s *loanUsecaseTestSuite) TestListUserRepaymentsSuccess() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorGetLoan() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorOwnership() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorApproval() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorGetRepayment() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorAlreadyPaid() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorInsufficientAmount() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorUpdateRepayment() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorCheckNextRepayment() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentErrorUpdateLoanStatus() {}

func (s *loanUsecaseTestSuite) TestUserSubmitRepaymentSuccess() {}
