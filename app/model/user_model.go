package model

import (
	"gorm.io/gorm"
	"time"
)

type UserResponse struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Username  string          `json:"username,omitempty"`
	Email     string          `json:"email,omitempty"`
	Phone     *string         `json:"phone,omitempty"`
	Photo     *string         `json:"photo,omitempty"`
	Token     *string         `json:"token,omitempty"`
	Role      string          `json:"role,omitempty"`
	Employees *[]UserResponse `json:"employees,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=255"`
}

type RegisterUserRequest struct {
	Name     string  `validate:"required,max=255" json:"name"`
	Username string  `validate:"required,max=255" json:"username"`
	Password string  `validate:"required,max=255" json:"password"`
	Email    string  `validate:"required,max=255" json:"email"`
	Phone    string  `validate:"max=16" json:"phone"`
	Role     string  `validate:"max=255" json:"role"`
	UserId   *string `json:"user_id,omitempty"`
}

type UpdateUserRequest struct {
	ID       string `validate:"required,max=36" json:"id"`
	Name     string `validate:"max=255" json:"name"`
	Email    string `validate:"max=255" json:"email"`
	Phone    string `validate:"max=16" json:"phone"`
	Username string `validate:"max=255" json:"username"`
	Password string `validate:"max=255" json:"password"`
}

type LoginUserRequest struct {
	Username string `validate:"required,max=255" json:"username"`
	Password string `validate:"required,max=255" json:"password"`
}

type LogoutUserRequest struct {
	ID string `validate:"required,max=36" json:"id"`
}

type GetUserRequest struct {
	ID string `validate:"required,max=36" json:"id"`
}
