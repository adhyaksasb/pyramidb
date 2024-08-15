package model

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"type:varchar(255);unique;not null"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	StarRailUID  uint   `gorm:"null"`
	Achievements string `gorm:"type:text;null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}