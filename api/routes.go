// task-IAM-service/api/routes.go

package api

import (
	"task-profile-service/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/dashboard", handlers.GetUserDashboard)
	router.GET("/profile", handlers.GetUserProfile)
	router.PUT("/profile/email", handlers.UpdateUserEmail)
	router.PUT("/profile/password", handlers.UpdateUserPassword)

}
