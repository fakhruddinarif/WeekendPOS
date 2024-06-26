package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:        product.ID,
		SKU:       product.SKU,
		Name:      product.Name,
		Category:  product.Category.Name,
		BuyPrice:  product.BuyPrice,
		SellPrice: product.SellPrice,
		Stock:     product.Stock,
		Photo:     product.Photo,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ProductToEvent(product *entity.Product) *model.ProductEvent {
	return &model.ProductEvent{
		ID:        product.ID,
		SKU:       product.SKU,
		Name:      product.Name,
		Category:  product.Category.Name,
		BuyPrice:  product.BuyPrice,
		SellPrice: product.SellPrice,
		Stock:     product.Stock,
		Photo:     product.Photo,
		User:      product.User.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
