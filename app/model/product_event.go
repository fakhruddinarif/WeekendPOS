package model

import "time"

type ProductEvent struct {
	ID         string    `json:"id,omitempty"`
	SKU        string    `json:"sku,omitempty"`
	Name       string    `json:"name,omitempty"`
	CategoryID string    `json:"category_id,omitempty"`
	BuyPrice   float64   `json:"buy_price,omitempty"`
	SellPrice  float64   `json:"sell_price,omitempty"`
	Stock      int       `json:"stock,omitempty"`
	Photo      string    `json:"photo,omitempty"`
	UserID     string    `json:"user_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (c *ProductEvent) GetId() string {
	return c.ID
}
