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

func (s *loanUsecaseTestSuite) TestSubmitLoanErrorStoreLoan() {}

func (s *loanUsecaseTestSuite) TestSubmitLoanErrorBulkStoreRepayment() {}

func (s *loanUsecaseTestSuite) TestSubmitLoanSuccess() {}

func (s *loanUsecaseTestSuite) TestListUserLoans() {}

func (s *loanUsecaseTestSuite) TestUserGetLoanErrorRepo() {}

func (s *loanUsecaseTestSuite) TestUserGetLoanErrorOwnership() {}

func (s *loanUsecaseTestSuite) TestUserGetLoanSuccess() {}

func (s *loanUsecaseTestSuite) TestListLoans() {}

func (s *loanUsecaseTestSuite) TestApproveLoan() {}

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
