package model

import "time"

type CategoryEvent struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (c *CategoryEvent) GetId() string {
	return c.ID
}
