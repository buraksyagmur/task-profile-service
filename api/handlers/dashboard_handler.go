// task-profile-service/api/handlers/dashboard_handler.go
package handlers

import (
	"fmt"
	"net/http"
	"task-profile-service/authentication"
	"task-profile-service/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Dashboard
// @Description Get the dashboard of the logged-in user.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ResponseOK "User dashboard retrieved successfully"
// @Failure 401 {object} models.ResponseError "Unauthorized"
// @Router /dashboard [get]
func GetUserDashboard(c *gin.Context) {
	fmt.Println(c.Request.Method)
	if c.Request.Method != "GET" {
		return
	}
	fmt.Println("auth")
	authentication.AuthenticateJWT()(c)
	username, _ := c.Get("username")
	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ResponseError{Error: "Failed to convert username to string"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome, " + usernameStr + "!",
	})
}
