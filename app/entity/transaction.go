package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID         string    `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	Customer   string    `gorm:"column:customer;type:varchar(255);not null"`
	Date       string    `gorm:"column:date;type:datetime;not null"`
	Total      float64   `gorm:"column:total;type:decimal(10,2);not null;"`
	User       User      `gorm:"foreignKey:user_id;references:id"`
	UserId     string    `gorm:"column:user_id;type:char(36);not null"`
	Employee   Employee  `gorm:"foreignKey:employee_id;references:id"`
	EmployeeId *string   `gorm:"column:employee_id;type:char(36);not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
