package middlewares

import (
	"mygram/models"
	"mygram/database"
	"net/http"
	"strconv"
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})

			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "Unauthenticated",
				"message": "you are nor allowed to access this data",
			})

			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})

			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "Unauthenticated",
				"message": "you are nor allowed to access this data",
			})

			return
		}

		c.Next()
	}
}

func SocialmediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialmediaId, err := strconv.Atoi(c.Param("socialmediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"error": "Bad Request",
				"message": "invalid parameter",
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Socialmedia := models.Socialmedia{}

		err = db.Select("user_id").First(&Socialmedia, uint(socialmediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})

			return
		}

		if Socialmedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "Unauthenticated",
				"message": "you are nor allowed to access this data",
			})

			return
		}

		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userIDStr := c.Param("userID")
		log.Println("User ID:", userIDStr) // Log user ID parameter

		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			log.Println("Error parsing user ID:", err) // Log parsing error
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid user ID parameter",
			})
			return
		}

		// Perform database query to check if user exists or has specific permissions
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			log.Println("Error fetching user from database:", err) // Log database query error
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "User not found",
			})
			return
		}

		// Optionally, check if the authenticated user has permission to access this user's data
		userData := c.MustGet("userData").(jwt.MapClaims)
		authenticatedUserID := uint(userData["id"].(float64))
		if authenticatedUserID != uint(userID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
