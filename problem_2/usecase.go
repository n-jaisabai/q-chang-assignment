package problem_2

import (
	"context"

	"github.com/n-jaisabai/q-chang-assignment/models"
)

// Usecase represent the problem_2 usecases
type Usecase interface {
	MakeChange(ctx context.Context, amount float64, productPrice float64) ([]*models.Money, error)
}
