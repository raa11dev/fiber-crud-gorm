package response

import "time"

type UserResponse struct {
	NIK       uint      `json:"nik" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email"`
	Password  string    `json:"-" gorm:"column:password"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
