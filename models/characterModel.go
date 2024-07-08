package model

type Character struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"size:255"`
	Tag            string `gorm:"size:255"`
	Rarity         int
	Element        string `gorm:"size:255"`
	PathID         uint   // Foreign key
	MaxSP          int
	ReleaseVersion string        `gorm:"size:255"`
	Icon           string        `gorm:"size:255"`
	Preview        string        `gorm:"size:255"`
	Portrait       string        `gorm:"size:255"`
	Path           Path          `gorm:"foreignKey:PathID"`
	CharacterStat  CharacterStat `gorm:"foreignKey:ID"`
}