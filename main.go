package main

import (
	"github.com/adhyaksasb/pyramidb/controllers"
	"github.com/adhyaksasb/pyramidb/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Write Service
	r.POST("/posts", controllers.CreatePosts)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Read Service
	r.GET("/posts", controllers.IndexPosts)
	r.GET("/characters", controllers.GetAllCharacters)
	r.GET("/characters/:tag", controllers.GetCharacterByTag)
	r.GET("/character-skills/:id", controllers.GetCharacterSkill)
	r.GET("/posts/:id", controllers.ShowPost)
	r.GET("/paths", controllers.GetAllPaths)
	r.GET("/elements", controllers.GetAllElements)
	r.GET("/achievements", controllers.GetAllAchievements)

	r.Run() // listen and serve on 0.0.0.0:8080
}
