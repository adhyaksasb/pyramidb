package model

type Element struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Desc  string `gorm:"type:text"`
	Color string `gorm:"size:255"`
	Icon  string `gorm:"size:255"`
}
