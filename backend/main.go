package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB environmentals
const (
	dbHost     = "atlas-db.c3ae4mu4i6yu.us-east-2.rds.amazonaws.com" // RDS endpoint on AWS
	dbUser     = "atlas_admin"
	dbPassword = "atlas_password"
	dbName     = "atlas"
	dbPort     = "5432"    // Default PostgreSQL port
	dbSSLMode  = "require" // SSL/TLS encryption between backend and db.
)

// Gorm will parase the CamelCase names and create snake_case labels.
type URL struct {
	ID          uint   `gorm:"primaryKey"`
	SatiPath    string `gorm:"unique;not null"`
	OriginalURL string `gorm:"not null"`
	CreatedAt   time.Time
}

func main() {
	// DSB: Database Source Name; configuration string.
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
		" dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode

	// 1. Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Returns *gorm.db
	if err != nil {
		panic("Failed to connect to database")
	}
	// 2. Automatically create/update SQL schema based on provided URL struct.
	db.AutoMigrate(&URL{})
	// 3. Initialize
	r := gin.Default()
	// 4-1. When POST requset made to /shorten endpoint,
	// shortenURL is called passing gin context (c) and *gorm.db instance (db)
	r.POST("/shorten", func(c *gin.Context) { shortenURL(c, db) })
	// 4-2. When GET request is made to "/<value>" route, the <value> is captured as SatiPath
	// under c *gin.Context, and retreivable as c.Param("SatiPath")
	r.GET("/:SatiPath", func(c *gin.Context) { redirectURL(c, db) })
	r.Run(":8080")
}

func shortenURL(c *gin.Context, db *gorm.DB) {
	var req struct {
		OriginalURL string `json:"original_url"`
		SatiPath    string `json:"sati_path"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // code: 400
		return
	}

	if req.OriginalURL != "" && req.SatiPath != "" {
		var existingShortURL URL
		// 1. Check for duplicate entry
		if err := db.Select("sati_path").Where("sati_path = ?", req.SatiPath).First(&existingShortURL).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "sati path already in use"}) // code: 409
			return
		}
		// 2. Create entry
		if err := db.Create(&URL{OriginalURL: req.OriginalURL, SatiPath: req.SatiPath}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create URL on DB"}) // code: 500
			return
		}
		c.JSON(http.StatusOK, gin.H{"sati_path": "http://localhost:8080/" + req.SatiPath}) // code: 200

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both original url and sati path must be provided"}) // code: 400
	}
}

func redirectURL(c *gin.Context, db *gorm.DB) {
	var url URL
	if err := db.Where("sati_path = ?", c.Param("SatiPath")).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
