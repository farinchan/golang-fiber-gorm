package models

import "time"

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Username  *string
	Email     string
	Password  string
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
