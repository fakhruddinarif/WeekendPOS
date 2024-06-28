package model

import "time"

type EmployeeResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateEmployeeRequest struct {
	UserID   string `validate:"required,max=36" json:"user_id"`
	Name     string `validate:"required,max=255" json:"name"`
	Email    string `validate:"required,max=255" json:"email"`
	Username string `validate:"required,max=255" json:"username"`
	Password string `validate:"required,max=255" json:"password"`
	Phone    string `validate:"max=16" json:"phone"`
	Address  string `validate:"max=255" json:"address"`
	Photo    string `validate:"max=255" json:"photo"`
}

type UpdateEmployeeRequest struct {
	UserID   string `validate:"required,max=36" json:"user_id"`
	ID       string `validate:"required,max=36" json:"id"`
	Name     string `validate:"required,max=255" json:"name"`
	Email    string `validate:"required,max=255" json:"email"`
	Username string `validate:"required,max=255" json:"username"`
	Password string `validate:"required,max=255" json:"password"`
	Phone    string `validate:"max=16" json:"phone"`
	Address  string `validate:"max=255" json:"address"`
	Photo    string `validate:"max=255" json:"photo"`
}

type ListEmployeeRequest struct {
	UserID   string `validate:"required,max=36" json:"user_id"`
	Name     string `validate:"max=255" json:"name"`
	Email    string `validate:"max=255" json:"email"`
	Username string `validate:"max=255" json:"username"`
	Page     int    `validate:"min=1" json:"page"`
	Size     int    `validate:"min=1,max=100" json:"size"`
}

type GetEmployeeRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}

type DeleteEmployeeRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}
