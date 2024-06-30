package model

type DetailTransactionResponse struct {
	ID      int     `json:"id,omitempty"`
	Amount  int     `json:"amount,omitempty"`
	Price   float64 `json:"price,omitempty"`
	Product string  `json:"product,omitempty"`
}

type CreateDetailTransactionRequest struct {
	Amount    int    `validate:"required" json:"amount"`
	ProductID string `validate:"required" json:"product_id"`
}
