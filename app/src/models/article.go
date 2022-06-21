package models

import "time"

//article構造体
type Article struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Init() {
	db := DatabaseConnection()
	db.AutoMigrate(&Article{})
}
