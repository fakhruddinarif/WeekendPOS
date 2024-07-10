package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func TransactionToResponse(transaction *entity.Transaction, detail []entity.DetailTransaction) *model.TransactionResponse {
	detailResponse := make([]model.DetailTransactionResponse, 0)
	for _, d := range detail {
		detailResponse = append(detailResponse, model.DetailTransactionResponse{
			ID:      d.ID,
			Product: d.Product.Name,
			Amount:  d.Amount,
			Price:   d.Price,
		})
	}

	return &model.TransactionResponse{
		ID:        transaction.ID,
		Customer:  transaction.Customer,
		Date:      transaction.Date,
		Total:     transaction.Total,
		Employee:  transaction.Employee.Name,
		Products:  detailResponse,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
}

func TransactionToEvent(transaction *entity.Transaction) *model.TransactionEvent {
	return &model.TransactionEvent{
		ID:         transaction.ID,
		Customer:   transaction.Customer,
		Date:       transaction.Date,
		Total:      transaction.Total,
		UserID:     transaction.UserId,
		EmployeeID: transaction.EmployeeId,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}
}
