package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/n-jaisabai/q-chang-assignment/models"
	"github.com/n-jaisabai/q-chang-assignment/problem_2"
)

type problem2Usecase struct {
	problem2Repo   problem_2.Repository
	contextTimeout time.Duration
}

// NewProblem2Usecase will create new an problem2Usecase object representation of problem_2.Usecase interface
func NewProblem2Usecase(repo problem_2.Repository, timeout time.Duration) problem_2.Usecase {
	return &problem2Usecase{
		problem2Repo:   repo,
		contextTimeout: timeout,
	}
}

// makeChange implements problem_2.Usecase
func (ucase *problem2Usecase) MakeChange(ctx context.Context, amount float64, productPrice float64) ([]*models.Money, error) {

	if amount < productPrice {
		return nil, errors.New("not enough money")
	}

	if amount == productPrice {
		return nil, errors.New("no change")
	}

	coins, err := ucase.problem2Repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var changeAmount = amount - productPrice
	var moneyChange = []*models.Money{}

	for _, coin := range coins {
		var quantity = changeAmount / coin.CoinType

		if quantity > float64(coin.CoinAmount) {
			if int(coin.CoinAmount) <= 0 {
				continue
			}

			moneyChange = append(moneyChange, &models.Money{
				Type:   coin.CoinType,
				Amount: coin.CoinAmount,
			})
			changeAmount -= coin.CoinType * float64(coin.CoinAmount)
		} else {
			if int(quantity) <= 0 {
				continue
			}

			moneyChange = append(moneyChange, &models.Money{
				Type:   coin.CoinType,
				Amount: int(quantity),
			})
			changeAmount -= coin.CoinType * float64(int(quantity))
		}
	}

	if changeAmount > 0 {
		return nil, errors.New("can't change money")
	}

	for _, money := range moneyChange {
		if err := ucase.problem2Repo.UpdateAmountByType(ctx, &models.CashierDesk{
			CoinType:   money.Type,
			CoinAmount: -money.Amount,
		}); err != nil {
			return nil, err
		}
	}

	return moneyChange, nil
}
