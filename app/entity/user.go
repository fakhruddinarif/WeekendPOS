package entity

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Photo     string    `gorm:"column:photo;type:varchar(255);not null"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	Username  string    `gorm:"column:username;type:varchar(255);not null;unique"`
	Password  string    `gorm:"column:password;type:varchar(255);not null"`
	Email     string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Phone     string    `gorm:"column:phone;type:varchar(16);not null"`
	Token     string    `gorm:"column:token;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli;autoCreateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
