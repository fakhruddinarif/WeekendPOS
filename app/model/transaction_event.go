package model

import "time"

type TransactionEvent struct {
	ID        string    `json:"id,omitempty"`
	Customer  string    `json:"customer,omitempty"`
	Date      string    `json:"date,omitempty"`
	Total     float64   `json:"total,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (t *TransactionEvent) GetId() string {
	return t.ID
}
