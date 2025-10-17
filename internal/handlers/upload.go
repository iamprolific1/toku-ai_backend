package handlers

import (
	"time"
	"os"
	"path/filepath"
	"net/http"
	"fmt"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"tokuai/internal/models"
)

func UploadHandler(db *gorm.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		// Get userID from context - the one set in AuthMiddleware
		userID, exists := c.Get("userID")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		//Get Upload File From Form
		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		// Uploads Directory
		uploadsDir := "/uploads"
		os.MkdirAll(uploadsDir, os.ModePerm)

		// Save file
		filePath := filepath.Join(uploadsDir, fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename))
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// create DB record
		upload := models.Upload{
			UserID: userID.(uint),
			FilePath: filePath,
			Status: "Pending",
			Transcript: "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&upload).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save upload record"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"upload":  upload,
		})
	}
}