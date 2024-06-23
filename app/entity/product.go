package entity

import "time"

type Product struct {
	ID        string    `gorm:"column:id;primaryKey;type:char(36);not null;unique;index"`
	SKU       string    `gorm:"column:sku;type:varchar(50);not null;"`
	Name      string    `gorm:"column:name;type:varchar(255);not null;"`
	Category  Category  `gorm:"foreignKey:category_id;references:id"`
	BuyPrice  float64   `gorm:"column:buy_price;type:decimal(10,2);not null;"`
	SellPrice float64   `gorm:"column:sell_price;type:decimal(10,2);not null;"`
	Stock     int       `gorm:"column:stock;type:int;not null;"`
	Photo     string    `gorm:"column:photo;type:varchar(255);null"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (p *Product) TableName() string {
	return "products"
}
