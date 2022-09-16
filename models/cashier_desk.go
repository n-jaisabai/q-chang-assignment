package models

import (
	"time"

	"github.com/google/uuid"
)

// CashierDesk represent the cashier desk model
type CashierDesk struct {
	ID         uuid.UUID `json:"id"`
	CoinType   float64   `json:"coin_type"`
	CoinAmount int       `json:"coin_amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
