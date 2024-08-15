package controllers

import (
	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"
	"github.com/gin-gonic/gin"
)

func GetAllAchievements(c *gin.Context) {
	// Get the achievements with their associated paths in a single query
	var achievements []model.Achievement
	initializers.DB.Preload("AchievementSeries").Order("title asc").Find(&achievements)

	var mergedAchievements []gin.H

	// Decode and store JSON data into the slice
	for _, achievement := range achievements {
		mergedAchievement := gin.H{
			"ID":              achievement.ID,
			"Series":          achievement.AchievementSeries.Title,
			"RelationID":             achievement.RelationID,
			"Title": achievement.Title,
			"Desc": achievement.Desc,
			"Hide": achievement.Hide,
			"Rarity":          achievement.Rarity,
			"Reward": achievement.Reward,
			"Version": achievement.Version,
		}
		mergedAchievements = append(mergedAchievements, mergedAchievement)
	}

	// Return it
	c.JSON(200, gin.H{
		"achievements": mergedAchievements,
	})
}

func GetAllAchievementSeries (c *gin.Context) {
	// Get the achievementSeries
	var achievementSeries []model.AchievementSeries
	initializers.DB.Find(&achievementSeries)

	// Return it
	c.JSON(200, gin.H{
		"achievement_series": achievementSeries,
	})
}