package model

import "time"

type UserEvent struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Username  string    `json:"username,omitempty"`
	Photo     *string   `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return u.ID
}
