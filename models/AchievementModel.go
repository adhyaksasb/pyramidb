package model

type Achievement struct {
	ID         uint   `gorm:"primaryKey"`
	SeriesID   uint   `gorm:"not null"`
	RelationID uint   `gorm:"not null"`
	Title      string `gorm:"type:varchar(255);not null"`
	Desc       string `gorm:"type:text;not null"`
	Hide       bool   `gorm:"not null"`
	Rarity     string `gorm:"type:varchar(255);not null"`
	Reward     int    `gorm:"not null"`
	Version    string `gorm:"type:varchar(255);not null"`

	// Define the foreign key relationship
	AchievementSeries AchievementSeries `gorm:"foreignKey:SeriesID;references:ID"`
}