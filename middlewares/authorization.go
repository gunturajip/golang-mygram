package middlewares

import (
	"golang-mygram/database"
	"golang-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid photo parameter",
			})
			return
		}

		Photo := models.Photo{}
		err = db.First(&Photo, uint(photoID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "photo doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Photo.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update / delete this photo",
				})
				return
			}
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(c.Param("commentID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid comment parameter",
			})
			return
		}

		Comment := models.Comment{}
		err = db.First(&Comment, uint(commentID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "comment doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Comment.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update / delete this comment",
				})
				return
			}
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid social media parameter",
			})
			return
		}

		Socialmedia := models.Socialmedia{}
		err = db.First(&Socialmedia, uint(socialMediaID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "social media doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Socialmedia.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update / delete this social media",
				})
				return
			}
		}
		c.Next()
	}
}

// func AdminAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		admin := userData["admin"].(bool)
// 		if !admin {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "you are not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }
