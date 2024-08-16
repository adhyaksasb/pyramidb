package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Username	string
		Email		string
		Password	string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hashed password",
		})

		return
	}

	// Create the user
	user := model.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create user",
	})
}

func Login(c *gin.Context) {
	// Get the email and password from request body
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up the requested user
	var user model.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username/email or password",
		})
		return
	}

	// Compare the sent password with the saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username/email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Determine the domain based on the request origin
	domain := ""
	if c.Request.Host == "pyramidb-fe.vercel.app" {
		domain = "pyramidb-fe.vercel.app"
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",   // Cookie name
		tokenString,       // Cookie value (JWT token)
		3600*24*7,         // Expiry (7 days)
		"/",               // Path
		domain,            // Domain (empty string for localhost)
		domain != "",      // Secure flag (true for non-localhost)
		true,              // HttpOnly flag
	)

	// Send token back in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Success to login",
	})
}


func Me(c *gin.Context) {
	user, _ := c.Get("user");
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged In",
		"user": user,
	})
}