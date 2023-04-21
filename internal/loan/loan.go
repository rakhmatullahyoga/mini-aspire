package loan

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/rakhmatullahyoga/mini-aspire/commons"
)

const (
	pageSize  = 10
	precision = 2
)

var (
	errLoanOwnership   = errors.New("invalid resource ownership")
	errLoanNotApproved = errors.New("your loan is not approved")
	errLoanAlreadyPaid = errors.New("your loan is already paid")
	errInsufficient    = errors.New("insufficient payment amount")
)

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "pending"
	LoanStatusApproved LoanStatus = "approved"
	LoanStatusPaid     LoanStatus = "paid"
)

type Loan struct {
	ID       string     `json:"id"`
	UserID   string     `json:"user_id"`
	Amount   int        `json:"amount"`
	Term     int        `json:"term"`
	LoanDate time.Time  `json:"loan_date"`
	Status   LoanStatus `json:"status"`
}

type RepaymentStatus string

const (
	RepaymentStatusPending RepaymentStatus = "pending"
	RepaymentStatusPaid    RepaymentStatus = "paid"
)

type Repayment struct {
	ID      string          `json:"id"`
	LoanID  string          `json:"loan_id"`
	Amount  float64         `json:"amount"`
	DueDate time.Time       `json:"due_date"`
	Status  RepaymentStatus `json:"status"`
}

//go:generate mockery --name=ILoanRepository --structname ILoanRepository --filename=ILoanRepository.go --output=mocks
type ILoanRepository interface {
	StoreLoan(loan *Loan) error
	ListUserLoans(userID string, offset, count int) ([]Loan, error)
	GetLoanByID(loanID string) (*Loan, error)
	ListLoans(offset, count int) ([]Loan, error)
	SetLoanStatus(loanID string, status LoanStatus) (*Loan, error)
}

//go:generate mockery --name=IRepaymentRepository --structname IRepaymentRepository --filename=IRepaymentRepository.go --output=mocks
type IRepaymentRepository interface {
	BulkStoreRepayment(repayments []Repayment) error
	ListLoanRepayments(loanID string, offset, count int) ([]Repayment, error)
	GetFirstPendingRepayment(loanID string) (*Repayment, error)
	UpdateRepayment(repayment Repayment) error
}

type LoanUsecase struct {
	loanRepo      ILoanRepository
	repaymentRepo IRepaymentRepository
}

func NewLoanUsecase(loanRepo ILoanRepository, repaymentRepo IRepaymentRepository) *LoanUsecase {
	return &LoanUsecase{
		loanRepo:      loanRepo,
		repaymentRepo: repaymentRepo,
	}
}

func (u *LoanUsecase) SubmitLoan(ctx context.Context, req LoanRequest) (*Loan, error) {
	userID := ctx.Value(commons.ClaimsKeyUserID).(string)
	date := req.LoanDate
	year, month, day := date.Date()
	date = time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	loan := Loan{
		UserID:   userID,
		Amount:   req.Amount,
		Term:     req.Term,
		LoanDate: date,
		Status:   LoanStatusPending,
	}
	err := u.loanRepo.StoreLoan(&loan)
	if err != nil {
		return nil, err
	}

	repayments := []Repayment{}
	dueDate := loan.LoanDate
	for i := 0; i < loan.Term; i++ {
		amount := (float64(loan.Amount) / float64(loan.Term))
		ratio := math.Pow(10, float64(precision))
		amount = math.Round(amount*ratio) / ratio
		dueDate = dueDate.AddDate(0, 0, 7)
		repayment := Repayment{
			LoanID:  loan.ID,
			Amount:  amount,
			DueDate: dueDate,
			Status:  RepaymentStatusPending,
		}
		repayments = append(repayments, repayment)
	}
	err = u.repaymentRepo.BulkStoreRepayment(repayments)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (u *LoanUsecase) ListUserLoans(ctx context.Context, page int) ([]Loan, error) {
	userID := ctx.Value(commons.ClaimsKeyUserID).(string)
	offset := (page - 1) * pageSize
	return u.loanRepo.ListUserLoans(userID, offset, pageSize)
}

func (u *LoanUsecase) UserGetLoan(ctx context.Context, loanID string) (*Loan, error) {
	userID := ctx.Value(commons.ClaimsKeyUserID).(string)
	loan, err := u.loanRepo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.UserID != userID {
		return nil, errLoanOwnership
	}

	return loan, nil
}

func (u *LoanUsecase) ListLoans(page int) ([]Loan, error) {
	offset := (page - 1) * pageSize
	return u.loanRepo.ListLoans(offset, pageSize)
}

func (u *LoanUsecase) ApproveLoan(loanID string) (*Loan, error) {
	return u.loanRepo.SetLoanStatus(loanID, LoanStatusApproved)
}

func (u *LoanUsecase) ListUserRepayments(ctx context.Context, loanID string, page int) ([]Repayment, error) {
	userID := ctx.Value(commons.ClaimsKeyUserID).(string)
	loan, err := u.loanRepo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.UserID != userID {
		return nil, errLoanOwnership
	}

	offset := (page - 1) * pageSize
	return u.repaymentRepo.ListLoanRepayments(loanID, offset, pageSize)
}

func (u *LoanUsecase) UserSubmitRepayment(ctx context.Context, loanID string, amount float64) (*Repayment, error) {
	userID := ctx.Value(commons.ClaimsKeyUserID).(string)
	loan, err := u.loanRepo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.UserID != userID {
		return nil, errLoanOwnership
	}

	if loan.Status != LoanStatusApproved {
		return nil, errLoanNotApproved
	}

	repayment, err := u.repaymentRepo.GetFirstPendingRepayment(loanID)
	if err != nil {
		return nil, err
	}

	if repayment != nil {
		if repayment.Amount > amount {
			return nil, errInsufficient
		}

		repayment.Status = RepaymentStatusPaid
		err = u.repaymentRepo.UpdateRepayment(*repayment)
		if err != nil {
			return nil, err
		}

		nextRepayment, err := u.repaymentRepo.GetFirstPendingRepayment(loanID)
		if err != nil && err != commons.ErrRecordNotFound {
			return nil, err
		}

		if nextRepayment == nil {
			_, err = u.loanRepo.SetLoanStatus(loanID, LoanStatusPaid)
			if err != nil {
				return nil, err
			}

			return repayment, nil
		} else {
			return repayment, nil
		}
	}

	return nil, errLoanAlreadyPaid
}
