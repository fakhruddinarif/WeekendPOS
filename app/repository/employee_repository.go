package repository

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	Repository[entity.Employee]
	Log *logrus.Logger
}

func NewEmployeeRepository(log *logrus.Logger) *EmployeeRepository {
	return &EmployeeRepository{
		Log: log,
	}
}

func (r *EmployeeRepository) Filter(request *model.ListEmployeeRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}
		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}
		if username := request.Username; username != "" {
			username = "%" + username + "%"
			tx = tx.Where("username LIKE ?", username)
		}
		if phone := request.Phone; phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}
		if address := request.Address; address != "" {
			address = "%" + address + "%"
			tx = tx.Where("address LIKE ?", address)
		}
		return tx
	}
}

func (r *EmployeeRepository) Search(db *gorm.DB, request *model.ListEmployeeRequest) ([]entity.Employee, int64, error) {
	var employees []entity.Employee
	var total int64 = 0

	if err := db.Scopes(r.Filter(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&employees).Error; err != nil {
		return nil, total, err
	}

	if err := db.Model(&entity.Employee{}).Scopes(r.Filter(request)).Count(&total).Error; err != nil {
		return nil, total, err
	}

	return employees, total, nil
}
