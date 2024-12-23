package models

import "time"

type User struct {
	UserID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name         string `gorm:"column:user_name;type:varchar;not null"`
	HashPassword string `gorm:"column:user_hash_password;type:varchar;not null"`
	UserRole     string `gorm:"column:user_role;type:varchar;not null"`
}
type Users []User

func (User) TableName() string {
	return "users"
}

type Order struct {
	OrderID      uint      `gorm:"primaryKey;autoIncrement;not null"`
	DeliveryType string    `gorm:"type:varchar;not null"`
	DeliveryTime time.Time `gorm:"type:timestamp;not null"`
	OrderTime    time.Time `gorm:"type:timestamp;not null"`
	TotalPrice   float64   `gorm:"type:float8;not null"`
	Canceled     bool      `gorm:"type:bool;not null"`
}
type Orders []Order

func (Order) TableName() string {
	return "orders"
}
