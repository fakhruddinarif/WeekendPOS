package entity

type DetailTransaction struct {
	ID          int         `gorm:"column:id;primaryKey;autoIncrement"`
	Amount      int         `gorm:"column:amount;type:int;not null"`
	Price       float64     `gorm:"column:price;type:decimal(10,2);not null"`
	Transaction Transaction `gorm:"foreignKey:transaction_id;references:id"`
	Product     Product     `gorm:"foreignKey:product_id;references:id"`
}

func (dt *DetailTransaction) TableName() string {
	return "detail_transactions"
}
