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
	initializers.DB.Find(&characters)

	// Return it
	c.JSON(200, gin.H{
		"characters": characters,
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