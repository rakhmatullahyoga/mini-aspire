package loan_test

import (
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/stretchr/testify/suite"
)

type repaymentRepositoryTestSuite struct {
	suite.Suite
	repo *loan.RepaymentRepository
}

func (s *repaymentRepositoryTestSuite) SetupTest() {
	s.repo = loan.NewRepaymentRepository()
}

func TestRepaymentRepository(t *testing.T) {
	suite.Run(t, new(repaymentRepositoryTestSuite))
}

func (s *repaymentRepositoryTestSuite) TestBulkStoreRepayment() {}

func (s *repaymentRepositoryTestSuite) TestListLoanRepaymentsNoData() {}

func (s *repaymentRepositoryTestSuite) TestListLoanRepaymentsFound() {}

func (s *repaymentRepositoryTestSuite) TestGetFirstPendingRepaymentNotFound() {}

func (s *repaymentRepositoryTestSuite) TestGetFirstPendingRepaymentFound() {}

func (s *repaymentRepositoryTestSuite) TestUpdateRepaymentErrorInvalid() {}

func (s *repaymentRepositoryTestSuite) TestUpdateRepaymentSuccess() {}
