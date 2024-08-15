package model

type Path struct {
	ID    uint   `gorm:"primaryKey"`
	Text  string `gorm:"size:255"`
	Name  string `gorm:"size:255"`
	Taunt int
	Desc  string `gorm:"type:text"`
	Icon  string `gorm:"size:255"`
}