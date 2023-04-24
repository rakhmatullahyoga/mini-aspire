package loan_test

import (
	"net/http"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan/mocks"
	"github.com/stretchr/testify/suite"
)

type userHandlerTestSuite struct {
	suite.Suite
	uc      *mocks.ILoanUsecase
	handler http.Handler
}

func (s *userHandlerTestSuite) SetupTest() {
	s.uc = new(mocks.ILoanUsecase)
	s.handler = loan.NewUserHandler(s.uc).Router()
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(userHandlerTestSuite))
}

func (s *userHandlerTestSuite) TestSubmitLoanInvalidRequestBody() {}

func (s *userHandlerTestSuite) TestSubmitLoanInvalidParameter() {}

func (s *userHandlerTestSuite) TestSubmitLoanErrorSubmitLoan() {}

func (s *userHandlerTestSuite) TestSubmitLoanSuccess() {}

func (s *userHandlerTestSuite) TestListUserLoansError() {}

func (s *userHandlerTestSuite) TestListUserLoansSuccess() {}

func (s *userHandlerTestSuite) TestUserGetLoanErrorOwnership() {}

func (s *userHandlerTestSuite) TestUserGetLoanErrorNotFound() {}

func (s *userHandlerTestSuite) TestUserGetLoanErrorUnhandled() {}

func (s *userHandlerTestSuite) TestUserGetLoanSuccess() {}

func (s *userHandlerTestSuite) TestUserGetLoanListUserRepaymentsErrorOwnership() {}

func (s *userHandlerTestSuite) TestUserGetLoanListUserRepaymentsErrorNotFound() {}

func (s *userHandlerTestSuite) TestUserGetLoanListUserRepaymentsErrorUnhandled() {}

func (s *userHandlerTestSuite) TestUserGetLoanListUserRepaymentsSuccess() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentInvalidRequestBody() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentInvalidParameter() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorOwnership() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorNotApproved() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorInsufficient() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorAlreadyPaid() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorNotFound() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentErrorUnhandled() {}

func (s *userHandlerTestSuite) TestUserSubmitRepaymentSuccess() {}
