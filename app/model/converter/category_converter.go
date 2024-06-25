package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func CategoryToEvent(category *entity.Category) *model.CategoryEvent {
	return &model.CategoryEvent{
		ID:        category.ID,
		Name:      category.Name,
		UserID:    category.UserId,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
