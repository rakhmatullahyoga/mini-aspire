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
