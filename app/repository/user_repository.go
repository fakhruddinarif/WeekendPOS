package repository

import (
	"WeekendPOS/app/entity"
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

func (r *Repository[T]) CountByUsername(db *gorm.DB, username string) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("username = ?", username).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindByUsername(db *gorm.DB, entity *T, username string) error {
	return db.Where("username = ?", username).Take(entity).Error
}
