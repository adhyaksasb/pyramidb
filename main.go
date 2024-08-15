package main

import (
	"github.com/adhyaksasb/pyramidb/controllers"
	"github.com/adhyaksasb/pyramidb/initializers"
	"github.com/adhyaksasb/pyramidb/middleware"

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
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.RequireAuth, controllers.Me)

	// Read Service
	r.GET("/posts", controllers.IndexPosts)
	r.GET("/characters",middleware.RequireAuth, controllers.GetAllCharacters)
	r.GET("/characters/:tag", controllers.GetCharacterByTag)
	r.GET("/character-skills/:id", controllers.GetCharacterSkill)
	r.GET("/posts/:id", controllers.ShowPost)
	r.GET("/paths", controllers.GetAllPaths)
	r.GET("/elements", controllers.GetAllElements)
	r.GET("/achievements", controllers.GetAllAchievements)
	r.GET("achievement-series", controllers.GetAllAchievementSeries)

	r.Run() // listen and serve on 0.0.0.0:8080
}
