package entity

import "time"

type Transaction struct {
	ID        string       `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	Customer  string       `gorm:"column:customer;type:varchar(255);not null"`
	Date      time.Weekday `gorm:"column:date;type:varchar(255);not null"`
	Total     float64      `gorm:"column:total;type:decimal(10,2);not null;"`
	User      User         `gorm:"foreignKey:user_id;references:id"`
	Employee  Employee     `gorm:"foreignKey:employee_id;references:id"`
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time    `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
