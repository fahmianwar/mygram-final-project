package controllers

import (
	"net/http"
	"strconv"

	"mygram-final-project/database"
	"mygram-final-project/helpers"
	"mygram-final-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetSocialMedias godoc
// @Summary Get details
// @Description Get details of all social media
// @Tags socialMedias
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socialMedias [get]
func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()
	socialMedias := []models.SocialMedia{}
	db.Find(&socialMedias)

	for i := 0; i < len(socialMedias); i++ {
		if err := db.Debug().Where("id = ?", socialMedias[i].UserID).First(&socialMedias[i].User).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMedias})
}

// GetSocialMedia godoc
// @Summary Get detail
// @Description Get detail of social media find by id
// @Tags socialMedias
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socialMedias/{socialMediaId} [get]
func GetSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}

	db := database.GetDB()
	if err := db.Where("id = ?", ctx.Param("socialMediaId")).First(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Debug().Where("id = ?", socialMedia.UserID).First(&socialMedia.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMedia})
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete of social media find by id
// @Tags socialMedias
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socialMedias/{socialMediaId} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialMedia := models.SocialMedia{}
	if err := db.Where("id = ?", ctx.Param("socialMediaId")).First(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!!"})
		return
	}

	db.Delete(&socialMedia)

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted SocialMedia Success"})
}

// CreateSocialMedia godoc
// @Summary Create the social media
// @Description Create the social media
// @Tags socialMedias
// @Accept json
// @Produce json
// @Param models.SocialMedia body models.SocialMedia true "create social media"
// @Success 200 {object} models.SocialMedia
// @Router /socialMedias [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)

	socialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = userID

	if err := db.Debug().Where("id = ?", userID).First(&socialMedia.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	err := db.Debug().Create(&socialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created SocialMedia Success", "data": socialMedia})
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)
	socialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = userID
	socialMedia.ID = uint(socialMediaId)

	err := db.Model(&socialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaUrl: socialMedia.SocialMediaUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMedia})
}
