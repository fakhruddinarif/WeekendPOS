package entity

import "time"

type Employee struct {
	ID           string        `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	Name         string        `gorm:"column:name;type:varchar(255);not null"`
	Email        string        `gorm:"column:email;type:varchar(255);not null;unique"`
	Username     string        `gorm:"column:username;type:varchar(255);not null;unique"`
	Password     string        `gorm:"column:password;type:varchar(255);not null"`
	Phone        string        `gorm:"column:phone;type:varchar(16);null"`
	Address      string        `gorm:"column:address;type:varchar(255);null"`
	Photo        string        `gorm:"column:photo;type:varchar(255);null"`
	User         User          `gorm:"foreignKey:user_id;references:id"`
	Transactions []Transaction `gorm:"foreignKey:employee_id;references:id"`
	CreatedAt    time.Time     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time     `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (e *Employee) TableName() string {
	return "employees"
}
