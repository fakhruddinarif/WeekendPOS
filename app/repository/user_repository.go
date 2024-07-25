package repository

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) CountByUsername(db *gorm.DB, username string) (int64, error) {
	var total int64
	err := db.Model(&entity.User{}).Where("username = ?", username).Count(&total).Error
	return total, err
}

func (r *UserRepository) FindByUsername(db *gorm.DB, user *entity.User, username string) error {
	return db.Where("username = ?", username).Take(user).Error
}

func (r *UserRepository) FindEmployeesByUserId(db *gorm.DB, request *model.SearchEmployeeRequest) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64 = 0

	if err := db.Scopes(r.Filter(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&users).Error; err != nil {
		return nil, total, err
	}

	if err := db.Model(&entity.User{}).Scopes(r.Filter(request)).Count(&total).Error; err != nil {
		return nil, total, err
	}

	return users, total, nil
}

func (r *UserRepository) Filter(request *model.SearchEmployeeRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserID)
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}
		if username := request.Username; username != "" {
			username = "%" + username + "%"
			tx = tx.Where("username LIKE ?", username)
		}
		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}
		return tx
	}
}
