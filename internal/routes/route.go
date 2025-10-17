package routes

import (
	"log"
	"strings"
	"tokuai/internal/config"
	"tokuai/internal/handlers"
	"tokuai/internal/middleware"
	"tokuai/internal/repository"

	"github.com/gin-gonic/gin"
)

func Routes() {
	cfg := config.LoadConfig()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(corsMiddleware())

	authHandler := handlers.AuthHandler{DB: repository.DB}

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/uploads", handlers.UploadHandler(repository.DB))
	}

	

	r.Run(cfg.Port)
	log.Printf("Server running on %s\n", cfg.Port)
}

func corsMiddleware() gin.HandlerFunc {
	// define allowed origins
	originsString := "http://localhost:5173,http://localhost:3000,http://localhost:5174"

	var allowedOrigins []string

	if originsString != "" {
		// Split the origin strings into individual origins and store them in allowedOrigins slice
		allowedOrigins = strings.Split(originsString, ",")
	}

	// Return the actual middleware handler function
	return func(c *gin.Context) {
		// function to check if a given origin is allowed.
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return  false
		} 

		// Get the origin header from the header request
		origin := c.Request.Header.Get("origin")

		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			// if the origin is allowed, set CORS header in the response
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		// Handle preflight OPTIONS requests by aborting with status 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Call the next handler
		c.Next()
	}
}