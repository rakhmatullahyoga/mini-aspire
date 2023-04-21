package loan_test

import (
	"net/http"
	"testing"

	"github.com/rakhmatullahyoga/mini-aspire/internal/loan"
	"github.com/rakhmatullahyoga/mini-aspire/internal/loan/mocks"
	"github.com/stretchr/testify/suite"
)

type adminHandlerTestSuite struct {
	suite.Suite
	uc      *mocks.IAdminUsecase
	handler http.Handler
}

func (s *adminHandlerTestSuite) SetupTest() {
	s.uc = new(mocks.IAdminUsecase)
	s.handler = loan.NewAdminHandler(s.uc).Router()
}

func TestAdminHandler(t *testing.T) {
	suite.Run(t, new(adminHandlerTestSuite))
}
