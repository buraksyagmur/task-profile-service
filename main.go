// task-profile-service/main.go

package main

import (
	"fmt"
	"log"
	"task-profile-service/api"      // Update the import path
	"task-profile-service/database" // Update the import path
	"task-profile-service/config"
	_ "task-profile-service/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Profile Service
// @version 1.0
// @description Profile Service for handling user dashboard and profile pages.
// @termsOfService http://example.com/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8081
// @BasePath /
func main() {
	config.LoadConfig("./config.json")
	database.Connect()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))
	api.SetupRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := 8081
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
