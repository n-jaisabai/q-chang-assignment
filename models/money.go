package models

// Money represent the money model
type Money struct {
	Type   float64 `json:"type"`
	Amount int     `json:"coin_amount"`
}
