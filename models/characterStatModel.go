package model

type CharacterStat struct {
	ID       uint    `gorm:"primaryKey"`
	HP       int     // integer
	ATK      int     // integer
	DEF      int     // integer
	SPD      int     // integer
	CritRate float64 `gorm:"column:crit_rate"` // float
	CritDmg  float64 `gorm:"column:crit_dmg"`  // float
}