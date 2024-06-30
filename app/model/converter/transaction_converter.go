package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:        transaction.ID,
		Customer:  transaction.Customer,
		Date:      transaction.Date,
		Total:     transaction.Total,
		Employee:  transaction.Employee.Name,
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
