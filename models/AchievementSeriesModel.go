package model

type AchievementSeries struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255);not null"`
}