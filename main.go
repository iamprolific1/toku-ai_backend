package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"tokuai/internal/config"
	"tokuai/internal/repository"
	"tokuai/internal/models"
	"tokuai/internal/handlers"
	"tokuai/internal/middleware"
)

func main() {
	cfg := config.LoadConfig()

	repository.Connect()

	repository.DB.AutoMigrate(
		&models.User{},
		&models.Upload{},
		&models.Output{},
	)

	authHandler := handlers.AuthHandler{DB: repository.DB}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	//Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		r.POST("/uploads", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "placeholder message"}) })
	}

	r.Run(cfg.Port)
	log.Printf("Server running on %s\n", cfg.Port)
}