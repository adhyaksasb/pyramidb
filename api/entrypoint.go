package api

import (
	"net/http"

	"github.com/adhyaksasb/pyramidb/controllers"
	"github.com/adhyaksasb/pyramidb/initializers"
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

	// Read Service
	r.GET("/posts", controllers.IndexPosts)
	r.GET("/characters", controllers.GetAllCharacters)
	r.GET("/characters/:tag", controllers.GetCharacterByTag)
	r.GET("/character-skills/:id", controllers.GetCharacterSkill)
	r.GET("/posts/:id", controllers.ShowPost)
	r.GET("/paths", controllers.GetAllPaths)
	r.GET("/elements", controllers.GetAllElements)
}

func init() {
	initializers.ConnectToDB()
	app = gin.New()
	r := app.Group("/api")
	route(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
