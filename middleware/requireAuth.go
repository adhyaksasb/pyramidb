package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
func RequireAuth(c *gin.Context) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Validate Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		// If the token is not prefixed with "Bearer ", reject the request
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user by ID from the token's subject
		var user model.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user to the request context
		c.Set("user", user)

		// Continue with the request
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}