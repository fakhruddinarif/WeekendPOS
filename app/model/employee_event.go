package model

import "time"

type EmployeeEvent struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Token     string    `json:"token,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (e *EmployeeEvent) GetId() string {
	return e.ID
}
