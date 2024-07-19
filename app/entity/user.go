package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string        `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	Photo        *string       `gorm:"column:photo;type:varchar(255);null"`
	Name         string        `gorm:"column:name;type:varchar(255);not null"`
	Username     string        `gorm:"column:username;type:varchar(255);not null;unique"`
	Password     string        `gorm:"column:password;type:varchar(255);not null"`
	Email        string        `gorm:"column:email;type:varchar(255);not null;unique"`
	Phone        *string       `gorm:"column:phone;type:varchar(16);null"`
	Token        *string       `gorm:"column:token;type:varchar(255);null"`
	Categories   []Category    `gorm:"foreignKey:user_id;references:id"`
	Products     []Product     `gorm:"foreignKey:user_id;references:id"`
	Employees    []Employee    `gorm:"foreignKey:user_id;references:id"`
	Transactions []Transaction `gorm:"foreignKey:user_id;references:id"`
	CreatedAt    time.Time     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time     `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
