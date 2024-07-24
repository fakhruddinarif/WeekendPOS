package model

import (
	"gorm.io/gorm"
	"time"
)

type UserEvent struct {
	ID        string         `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Username  string         `json:"username,omitempty"`
	Photo     *string        `json:"photo,omitempty"`
	Role      string         `json:"role,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return u.ID
}
