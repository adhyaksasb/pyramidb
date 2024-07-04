package controllers

import (
	"encoding/json"
	"log"

	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"
	"github.com/gin-gonic/gin"
)

func GetAllCharacters(c *gin.Context) {
	// Get the posts
	var characters []model.Character
	initializers.DB.Order("name asc").Find(&characters)


	var mergedCharacters []gin.H

	var path model.Path
	// Decode and store JSON data into the slice
	for _, character := range characters {
		initializers.DB.Model(&path).First(&path, character.Path_id)

		mergedCharacter := gin.H{
			"ID": character.ID,
			"Name": character.Name,
			"Tag": character.Tag,
			"Rarity": character.Rarity,
			"Element": character.Element,
			"Path_name": path.Name,
			"Max_sp": character.Max_sp,
			"Taunt": path.Taunt,
			"Release_version": character.Release_version,
			"Icon": character.Icon,
			"Path_icon": path.Icon,
			"Preview": character.Preview,
			"Portrait": character.Portrait,
		}
		mergedCharacters = append(mergedCharacters, mergedCharacter)
	}

	// Return it
	c.JSON(200, gin.H{
		"characters": mergedCharacters,
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

    // Create a new slice to store decoded JSON data
    var decodedLevels [][]string

	var mergedSkills []gin.H

    // Decode and store JSON data into the slice
    for _, skill := range characterSkills {
        var levelData []string
		
        if err := json.Unmarshal([]byte(skill.Level), &levelData); err != nil {
            log.Printf("Failed to unmarshal JSON data: %v", err)
            continue
        }
		
		decodedLevels = append(decodedLevels, levelData)

		mergedSkill := gin.H{
			"id":           skill.ID,
			"character_id": skill.CharacterID,
			"name":         skill.Name,
			"max_level": 	skill.MaxLevel,
			"type": 		skill.Type,
			"effect": 		skill.Effect,
			"level": 		decodedLevels,
			"icon": 		skill.Icon,
		}
		mergedSkills = append(mergedSkills, mergedSkill)
    }

	// Return it
	c.JSON(200, gin.H{
		"skills": mergedSkills,
	})
}