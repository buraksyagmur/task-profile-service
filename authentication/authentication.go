// task-profile-service/authentication/authentication.go
package authentication

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenString := authHeader[len(prefix):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretkey123"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized3"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized4"})
			c.Abort()
			return
		}
		userID, ok := claims["user_id"].(float64)
		username, ok := claims["user_name"].(string)
		useremail, ok := claims["user_email"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized5"})
			c.Abort()
			return
		}
		c.Set("user_id", uint(userID))
		c.Set("username", username)
		c.Set("useremail", useremail)

		c.Next()
	}
}
