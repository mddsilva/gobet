package dao

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false;column:created_at"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at"`
}

type User struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	Name     string  `gorm:"type:varchar(100);not null"`
	Email    string  `gorm:"type:varchar(100);unique;not null"`
	Password string  `gorm:"type:varchar(255);not null"`
	Balance  float64 `gorm:"type:decimal(18,2);default:0.00"`
	BaseModel
}
