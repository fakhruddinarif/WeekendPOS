package model

import "time"

type CategoryResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateCategoryRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	Name   string `validate:"required,max=255" json:"name"`
}

type UpdateCategoryRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
	Name   string `validate:"required,max=255" json:"name"`
}

type SearchCategoryRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	Name   string `validate:"max=255" json:"name"`
	Page   int    `validate:"min=1" json:"page"`
	Size   int    `validate:"min=1,max=100" json:"size"`
}

type GetCategoryRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}
type DeleteCategoryRequest struct {
	UserID string `validate:"required,max=36" json:"user_id"`
	ID     string `validate:"required,max=36" json:"id"`
}
