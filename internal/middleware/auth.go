package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"tokuai/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		// Set user ID in context
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}