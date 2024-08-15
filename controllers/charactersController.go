package controllers

import (
	"encoding/json"
	"log"

	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"
	"github.com/gin-gonic/gin"
)

func GetAllCharacters(c *gin.Context) {
	// Get the characters with their associated paths in a single query
	var characters []model.Character
	initializers.DB.Preload("Path").Order("name asc").Find(&characters)

	var mergedCharacters []gin.H

	// Decode and store JSON data into the slice
	for _, character := range characters {
		mergedCharacter := gin.H{
			"ID":              character.ID,
			"Name":            character.Name,
			"Tag":             character.Tag,
			"Rarity":          character.Rarity,
			"Element":         character.Element,
			"Path_name":       character.Path.Name,
			"Max_sp":          character.MaxSP,
			"Taunt":           character.Path.Taunt,
			"Release_version": character.ReleaseVersion,
			"Icon":            character.Icon,
			"Path_icon":       character.Path.Icon,
			"Preview":         character.Preview,
			"Portrait":        character.Portrait,
		}
		mergedCharacters = append(mergedCharacters, mergedCharacter)
	}

	// Return it
	c.JSON(200, gin.H{
		"characters": mergedCharacters,
	})
}

func GetCharacterByTag(c *gin.Context) {
	tag := c.Param("tag")

	// Get the characters with their associated paths in a single query
	var characters []model.Character
	initializers.DB.Preload("Path").Preload("CharacterStat").Where("tag = ?", tag).First(&characters)

	var mergedCharacters []gin.H

	// Decode and store JSON data into the slice
	for _, character := range characters {
		mergedCharacter := gin.H{
			"ID":              character.ID,
			"Name":            character.Name,
			"Tag":             character.Tag,
			"Rarity":          character.Rarity,
			"Element":         character.Element,
			"Max_sp":          character.MaxSP,
			"Release_version": character.ReleaseVersion,
			"Icon":            character.Icon,
			"Preview":         character.Preview,
			"Portrait":        character.Portrait,
			"Path": gin.H{
				"Path_name":       character.Path.Name,
				"Taunt":           character.Path.Taunt,
				"Path_icon":       character.Path.Icon,
			},
			"Status": gin.H{
				"HP": character.CharacterStat.HP,
				"ATK": character.CharacterStat.ATK,
				"DEF": character.CharacterStat.DEF,
				"SPD": character.CharacterStat.SPD,
				"Crit_rate": character.CharacterStat.CritRate,
				"Crit_dmg": character.CharacterStat.CritDmg,
			},
		}
		mergedCharacters = append(mergedCharacters, mergedCharacter)
	}

	// Return it
	c.JSON(200, gin.H{
		"character": mergedCharacters,
	})
}

func GetAllPaths(c *gin.Context) {
	var paths []model.Path
	initializers.DB.Order("name asc").Find(&paths)

	c.JSON(200, gin.H{
		"paths": paths,
	})
}

func GetAllElements(c *gin.Context) {
	var elements []model.Element
	initializers.DB.Order("name asc").Find(&elements)

	c.JSON(200, gin.H{
		"elements": elements,
	})
}

func GetCharacterSkill(c *gin.Context) {
	id := c.Param("id")

	// Get the posts
	var characterSkills []model.CharacterSkill
	initializers.DB.Where("character_id = ?", id).Find(&characterSkills)

	mergedSkills := make([]gin.H, len(characterSkills))

    // Decode and store JSON data into the slice
    for i, skill := range characterSkills {
		
        var decodedLevels []string

		// Unmarshal JSON data only if the skill type is not "Technique"
		if skill.Type != "Technique" {
			if err := json.Unmarshal([]byte(skill.Level), &decodedLevels); err != nil {
				log.Printf("Failed to unmarshal JSON data: %v", err)
				continue
			}
		}

		mergedSkills[i] = gin.H{
			"id":           skill.ID,
			"character_id": skill.CharacterID,
			"name":         skill.Name,
			"max_level":    skill.MaxLevel,
			"type":         skill.Type,
			"effect":       skill.Effect,
			"description": skill.Description,
			"level":        decodedLevels,
			"icon":         skill.Icon,
		}
    }

	c.JSON(200, gin.H{
		"skills": mergedSkills,
	})
}