package problem_2

import (
	"context"

	"github.com/n-jaisabai/q-chang-assignment/models"
)

// Repository represent the problem_2 repository contract
type Repository interface {
	Fetch(ctx context.Context) (res []*models.CashierDesk, err error)
	UpdateAmountByType(ctx context.Context, ar *models.CashierDesk) error
}
