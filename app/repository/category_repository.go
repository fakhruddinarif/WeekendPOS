package repository

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Repository[entity.Category]
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Log: log,
	}
}
func (r *CategoryRepository) FindByIdAndUserId(db *gorm.DB, category *entity.Category, id string, user string) error {
	return db.Where("id = ? AND user_id = ?", id, user).Take(category).Error
}

func (r *CategoryRepository) Search(db *gorm.DB, request *model.SearchCategoryRequest) ([]entity.Category, int64, error) {
	var categories []entity.Category
	var total int64 = 0

	if err := db.Scopes(r.Filter(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&categories).Error; err != nil {
		return nil, total, err
	}

	if err := db.Model(&entity.Category{}).Scopes(r.Filter(request)).Count(&total).Error; err != nil {
		return nil, total, err
	}

	return categories, total, nil
}

func (r *CategoryRepository) Filter(request *model.SearchCategoryRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserID)
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}
		return tx
	}
}
