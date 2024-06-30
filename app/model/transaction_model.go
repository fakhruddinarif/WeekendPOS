package model

import "time"

type TransactionResponse struct {
	ID        string                      `json:"id,omitempty"`
	Customer  string                      `json:"customer,omitempty"`
	Date      time.Time                   `json:"date,omitempty"`
	Total     float64                     `json:"total,omitempty"`
	Employee  string                      `json:"employee,omitempty"`
	Products  []DetailTransactionResponse `json:"products,omitempty"`
	CreatedAt time.Time                   `json:"created_at,omitempty"`
	UpdatedAt time.Time                   `json:"updated_at,omitempty"`
}

type CreateTransactionRequest struct {
	UserID     string                           `validate:"required,max=36" json:"user_id"`
	EmployeeID string                           `validate:"max=36" json:"employee_id"`
	Customer   string                           `validate:"required,max=255" json:"customer"`
	Date       time.Time                        `validate:"required" json:"date"`
	Products   []CreateDetailTransactionRequest `validate:"required,dive" json:"products"`
}

type GetTransactionRequest struct {
	UserID     string `validate:"required,max=36" json:"user_id"`
	EmployeeID string `validate:"max=36" json:"employee_id"`
	ID         string `validate:"required,max=36" json:"id"`
}

type ListTransactionRequest struct {
	UserID string    `validate:"required,max=36" json:"user_id"`
	Date   time.Time `validate:"max=255" json:"date"`
	Page   int       `validate:"min=1" json:"page"`
	Size   int       `validate:"min=1,max=100" json:"size"`
}

type UpdateTransactionRequest struct {
	UserID     string `validate:"required,max=36" json:"user_id"`
	EmployeeID string `validate:"max=36" json:"employee_id"`
	ID         string `validate:"required,max=36" json:"id"`
	Customer   string `validate:"required,max=255" json:"customer"`
	Date       string `validate:"required" json:"date"`
}
