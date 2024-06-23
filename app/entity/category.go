package entity

import "time"

type Category struct {
	ID        string    `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
	Products  []Product `gorm:"foreignKey:category_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (c *Category) TableName() string {
	return "categories"
}
