package repository

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (r *ProductRepository) FindById(db *gorm.DB, product *entity.Product, id string, user string) error {
	return db.Where("id = ? AND user_id = ?", id, user).Take(product).Error
}

func (r *ProductRepository) Filter(request *model.SearchProductRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserID)
		if sku := request.SKU; sku != "" {
			sku = "%" + sku + "%"
			tx = tx.Where("sku LIKE ?", sku)
		}
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}
		return tx
	}
}

func (r *ProductRepository) Search(db *gorm.DB, request *model.SearchProductRequest) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64 = 0

	if err := db.Scopes(r.Filter(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&products).Error; err != nil {
		return nil, total, err
	}

	if err := db.Model(&entity.Product{}).Scopes(r.Filter(request)).Count(&total).Error; err != nil {
		return nil, total, err
	}

	return products, total, nil
}
