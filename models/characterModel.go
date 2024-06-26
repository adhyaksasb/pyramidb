package model

type Character struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Tag             string
	Rarity          uint8
	Element         string
	Path_id         uint8
	Max_sp          uint16
	Release_version string
	Icon            string
	Preview         string
	Portrait        string
}