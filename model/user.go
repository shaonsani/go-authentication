package model

import "time"

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255)"`
	Gender    string `gorm:"type:varchar(10);default:other"` // Restricting length
	Mobile    string `gorm:"type:varchar(20);unique"`
	Address   string `gorm:"type:varchar(255)"`
	Photo     string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255);unique"`
	Password  string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
