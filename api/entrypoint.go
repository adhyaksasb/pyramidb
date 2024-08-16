package api

import (
	"net/http"
	"time"

	"github.com/adhyaksasb/pyramidb/controllers"
	"github.com/adhyaksasb/pyramidb/initializers"
	"github.com/adhyaksasb/pyramidb/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var(
	app *gin.Engine
)

func route(r *gin.RouterGroup) {
	// Write Service
	r.POST("/posts", controllers.CreatePosts)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.RequireAuth, controllers.Me)

	// Read Service
	r.GET("/posts", controllers.IndexPosts)
	r.GET("/characters", controllers.GetAllCharacters)
	r.GET("/characters/:tag", controllers.GetCharacterByTag)
	r.GET("/character-skills/:id", controllers.GetCharacterSkill)
	r.GET("/posts/:id", controllers.ShowPost)
	r.GET("/paths", controllers.GetAllPaths)
	r.GET("/elements", controllers.GetAllElements)
	r.GET("/achievements", controllers.GetAllAchievements)
	r.GET("/achievement-series", controllers.GetAllAchievementSeries)
}

func init() {
	initializers.ConnectToDB()
	app = gin.New()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://pyramidb-fe.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	app.Use(cors.New(config))
	r := app.Group("/api")
	route(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
