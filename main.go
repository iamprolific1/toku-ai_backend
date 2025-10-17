package main

import (

	"tokuai/internal/repository"
	"tokuai/internal/models"
	"tokuai/internal/routes"
	"tokuai/internal/config"
)

func main() {
	config.LoadConfig()

	repository.Connect()

	repository.DB.AutoMigrate(
		&models.User{},
		&models.Upload{},
		&models.Output{},
	)

	routes.Routes()


	// r.POST("/auth/register", authHandler.Register)
	// r.POST("/auth/login", authHandler.Login)

	//Protected routes
	

	
}