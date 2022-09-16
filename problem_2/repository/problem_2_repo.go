package repository

import (
	"context"

	"github.com/n-jaisabai/q-chang-assignment/models"
	"github.com/n-jaisabai/q-chang-assignment/problem_2"
	"gorm.io/gorm"
)

type sqlProblem2Repository struct {
	Conn *gorm.DB
}

func NewSqlProblem2Repository(Conn *gorm.DB) problem_2.Repository {
	return &sqlProblem2Repository{Conn}
}

func (r *sqlProblem2Repository) Fetch(ctx context.Context) (res []*models.CashierDesk, err error) {
	tx := r.Conn.WithContext(ctx)

	if err := tx.Order("coin_type desc").Find(&res).Error; err != nil {
		return nil, err
	}

	return
}

func (r *sqlProblem2Repository) UpdateAmountByType(ctx context.Context, m *models.CashierDesk) error {
	tx := r.Conn.WithContext(ctx)
	tx = tx.Model(&models.CashierDesk{})
	tx = tx.Where("coin_type = ?", m.CoinType)
	if err := tx.Update("coin_amount", gorm.Expr("coin_amount + ?", m.CoinAmount)).Error; err != nil {
		return err
	}

	return nil
}
