package model

type CharacterSkill struct {
	ID          uint   `gorm:"primaryKey"`
	CharacterID uint   `gorm:"column:character_id"`
	Name        string `gorm:"column:name"`
	MaxLevel    int    `gorm:"column:max_level"`
	Element     string `gorm:"column:element"`
	Type        string `gorm:"column:type"`
	Effect      string `gorm:"column:effect"`
	SimpleDesc  string `gorm:"column:simple_desc"`
	Description string `gorm:"column:description"`
	Level       []byte `gorm:"column:level"` // Assuming JSON is stored as []byte in Go
	Icon        string `gorm:"column:icon"`
}

func (CharacterSkill) TableName() string {
	return "character_skills"
}