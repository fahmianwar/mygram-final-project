package middlewares

import (
	"mygram-final-project/database"
	"mygram-final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorizad",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		err = db.Select("user_id").First(&photo, uint(photoId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorizad",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&socialMedia, uint(socialMediaId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if socialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorizad",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
