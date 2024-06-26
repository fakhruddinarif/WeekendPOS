package model

import "time"

type ProductResponse struct {
	ID        string    `json:"id,omitempty"`
	SKU       string    `json:"sku,omitempty"`
	Name      string    `json:"name,omitempty"`
	Category  string    `json:"category,omitempty"`
	BuyPrice  float64   `json:"buy_price,omitempty"`
	SellPrice float64   `json:"sell_price,omitempty"`
	Stock     int       `json:"stock,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	UserID     string    `validate:"required,max=36" json:"user_id"`
	SKU        string    `validate:"required,max=255" json:"sku"`
	Name       string    `validate:"required,max=255" json:"name"`
	CategoryID string    `validate:"required,max=36" json:"category_id"`
	BuyPrice   float64   `validate:"required" json:"buy_price"`
	SellPrice  float64   `validate:"required" json:"sell_price"`
	Stock      int       `validate:"required" json:"stock"`
	Photo      string    `json:"photo"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type UpdateProductRequest struct {
	UserID     string  `validate:"required,max=36" json:"user_id"`
	ID         string  `validate:"required,max=36" json:"id"`
	SKU        string  `validate:"required,max=255" json:"sku"`
	Name       string  `validate:"required,max=255" json:"name"`
	CategoryID string  `validate:"required,max=36" json:"category_id"`
	BuyPrice   float64 `validate:"required" json:"buy_price"`
	SellPrice  float64 `validate:"required" json:"sell_price"`
	Stock      int     `validate:"required" json:"stock"`
	Photo      string  `json:"photo"`
}

type SearchProductRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	SKU    string `validate:"max=255" json:"sku"`
	Name   string `validate:"max=255" json:"name"`
	Page   int    `validate:"min=1" json:"page"`
	Size   int    `validate:"min=1,max=100" json:"size"`
}

type GetProductRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}

type DeleteProductRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}
