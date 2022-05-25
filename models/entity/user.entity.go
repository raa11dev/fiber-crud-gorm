package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	NIK       uint           `json:"nik" gorm:"primaryKey"`
	Nama      string         `json:"nama"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
