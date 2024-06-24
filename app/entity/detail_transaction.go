package entity

type DetailTransaction struct {
	ID            int         `gorm:"column:id;primaryKey;autoIncrement"`
	Amount        int         `gorm:"column:amount;type:int;not null"`
	Price         float64     `gorm:"column:price;type:decimal(10,2);not null"`
	Transaction   Transaction `gorm:"foreignKey:transaction_id;references:id"`
	TransactionId string      `gorm:"column:transaction_id;type:char(36);not null"`
	Product       Product     `gorm:"foreignKey:product_id;references:id"`
	ProductId     string      `gorm:"column:product_id;type:char(36);not null"`
}

func (dt *DetailTransaction) TableName() string {
	return "detail_transactions"
}
