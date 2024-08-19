package controllers

import (
	"net/http"
	"strings"

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

func AddAchievementToUser(c *gin.Context) {
	// Extract the achievement ID from the URL
	achievementID := c.Param("achievementID")

	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Type assertion to access the user struct
	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve user information"})
		return
	}

	// Explode the Achievements string into a slice
	achievements := strings.Split(userModel.Achievements, ",")

	// Check if the achievement ID exists in the slice
	index := -1
	for i, id := range achievements {
		if id == achievementID {
			index = i
			break
		}
	}

	// If the achievement ID exists, remove it; if not, add it
	if index != -1 {
		// Remove the achievement
		achievements = append(achievements[:index], achievements[index+1:]...)
	} else {
		// Add the achievement
		achievements = append(achievements, achievementID)
	}

	// Implode the slice back into a comma-separated string
	userModel.Achievements = strings.Join(achievements, ",")

	// Update the user's achievements in the database
	if err := initializers.DB.Model(&userModel).Update("Achievements", userModel.Achievements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update achievements"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Add new achievements",
		"User": gin.H{
			"ID":        userModel.ID,
			"Username":  userModel.Username,
			"Email":     userModel.Email,
			"StarRailUID": userModel.StarRailUID,
			"Achievements": userModel.Achievements,
		},
	})
}
